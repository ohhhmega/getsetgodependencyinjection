package newrelic

import (
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
)

type Agent struct {
	a *newrelic.Application
}

func New(name, key string) (*Agent, error) {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(name),
		newrelic.ConfigLicense(key),
		//nrzap.ConfigLogger(logger.Logger),
	)
	if err != nil {
		return &Agent{}, errors.Wrap(err, "initiating newrelic session")
	}
	// Wait for the application to connect.
	if err = app.WaitForConnection(30 * time.Second); nil != err {
		return nil, err
	}
	if err != nil {
		return &Agent{}, errors.Wrap(err, "getting newrelic session")
	}
	return &Agent{app}, nil
}

func (ag Agent) StartTransaction(key string) *newrelic.Transaction {
	return ag.a.StartTransaction(key)
}

func (ag Agent) RecordCustomMetric(key string) {
	go ag.a.RecordCustomMetric(key, 1)
}

func (ag Agent) RecordCustomEvent(key string, params map[string]interface{}) {
	go ag.a.RecordCustomEvent(key, params)
}

func (ag Agent) NewRelicApplication() *newrelic.Application {
	return ag.a
}
