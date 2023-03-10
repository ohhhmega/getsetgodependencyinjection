package aws

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"bitbucket.org/junglee_games/getsetgo/common_errors"
	"bitbucket.org/junglee_games/getsetgo/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

type S3Store struct {
	uploader   *s3manager.Uploader
	s3Service  *s3.S3
	bucketName string
	jobQueue   chan *models.FileData
	//AckChan will receive a *models.FileData once its upload is completed by SaveAsync method
	ackChan    chan *models.FileData
	maxRetries int
	wg         *sync.WaitGroup
	logger     *log.Logger
}

func newS3Session() (*session.Session, error) {
	cred := aws.NewConfig()
	cred.WithRegion("ap-south-1")
	// Initialize a session in ap-south-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	return session.NewSession(cred)

}

func newS3Uploader(session *session.Session) (*s3manager.Uploader, error) {

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
	uploader := s3manager.NewUploader(session)
	uploader.MaxUploadParts = 1

	return uploader, nil
}

func NewS3Store(bucketName string, uploaderCount, maxRetries int, logger *log.Logger) (*S3Store, error) {

	var err error
	s3Store := S3Store{bucketName: bucketName, maxRetries: maxRetries}

	s3Store.jobQueue = make(chan *models.FileData, 100)
	s3Store.ackChan = make(chan *models.FileData, 100)

	s3Store.logger = logger

	session, err := newS3Session()
	if err != nil {
		return nil, err
	}

	s3Store.uploader, err = newS3Uploader(session)
	if err != nil {
		return nil, err
	}

	s3Store.s3Service = s3.New(session)

	s3Store.wg = new(sync.WaitGroup)
	for i := 0; i < uploaderCount; i++ {
		s3Store.wg.Add(1)
		go s3Store.startUploader()
	}

	return &s3Store, nil
}

func (s3S *S3Store) Save(filesData *models.FileData) (string, error) {
	ptr := strings.NewReader(filesData.Data)

	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	op, err := s3S.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3S.bucketName),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(filesData.Name),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: ptr,
	})
	if err != nil {
		err = errors.Wrap(err, "uploading to s3")
		return "", err
	}
	return op.Location, nil
}

// it will return signed url and error
func (s3S *S3Store) TemporarySave(filesData *models.FileData, expriryMinutes int) (string, error) {

	_, err := s3S.Save(filesData)
	if err != nil {
		return "", errors.Wrap(common_errors.ErrS3PresignFailed, err.Error())
	}
	return s3S.GetSignedURL(filesData.Name, expriryMinutes)
}

//it push the object into queue
func (s3S *S3Store) SaveAsync(filesData *models.FileData) {
	s3S.jobQueue <- filesData
}

func (s3S *S3Store) GetAckChan() chan *models.FileData {
	return s3S.ackChan
}

func (s3S *S3Store) StopUploading() {
	close(s3S.jobQueue)
}

func (s3S *S3Store) WaitForFinishingUpload() {
	s3S.wg.Wait()
	close(s3S.ackChan)
}

// it will take filesData from jobQueue and upload to s3 and ack by using ackqueue
func (s3S *S3Store) startUploader() {
	for filesData := range s3S.jobQueue {
		var path string
		var err error
		maxRetries := s3S.maxRetries
		for i := 1; i < maxRetries; i++ {
			path, err = s3S.Save(filesData)
			if err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		if err != nil {
			filesData.UploadInfo.Error = err
			s3S.ackChan <- filesData
			continue
		}
		filesData.UploadInfo.Path = path
		s3S.ackChan <- filesData
		continue
	}
	s3S.wg.Done()
	s3S.logger.Print("jobQueue of s3 uploader closed , exiting startUploader function ")
}

func (s3S *S3Store) DownloadFile(filename string) ([]byte, error) {

	// Define the parameters for the download request
	params := &s3.GetObjectInput{
		Bucket: aws.String(s3S.bucketName),
		Key:    aws.String(filename),
	}
	// Download the file
	resp, err := s3S.s3Service.GetObject(params)
	if err != nil {
		return nil, errors.Wrap(common_errors.ErrS3FileDownload, err.Error())
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(common_errors.ErrS3FileBodyReader, err.Error())
	}
	return bytes, nil
}

func (s3S *S3Store) GetSignedURL(filepath string, expriryMinutes int) (string, error) {

	// Define the parameters for the download request
	req, _ := s3S.s3Service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3S.bucketName),
		Key:    aws.String(filepath),
	})
	// Download the file
	url, err := req.Presign(time.Minute * time.Duration(expriryMinutes))
	if err != nil {
		return "", errors.Wrap(common_errors.ErrS3FileDownload, err.Error())
	}

	return url, nil
}
