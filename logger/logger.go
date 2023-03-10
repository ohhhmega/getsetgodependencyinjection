package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger    *zap.Logger
	appNameField zap.Field
	middleLayers []MiddleLayer
)

func init() {
	Config{AppName: "default", Build: "dev"}.InitiateLogger()
	appNameField = zap.Field{Key: "App", Type: zapcore.StringType, String: "default"}
	middleLayers = make([]MiddleLayer, 0)
}

func (c Config) InitiateLogger() error {
	var err error
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = timeKey
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var zapConfig zap.Config
	if c.Build == "prod" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}
	zapConfig.DisableStacktrace = true

	zapConfig.EncoderConfig = encoderConfig
	zapLogger, err = zapConfig.Build(zap.AddStacktrace(zapcore.ErrorLevel))

	appNameField.String = c.AppName
	return err
}

func AddMiddleLayers(middlelayers ...MiddleLayer) {
	middleLayers = append(middleLayers, middlelayers...)
}

func Info(ctx context.Context, format string, a ...any) {
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Info(msg, fields.fields...)
}

func Infow(ctx context.Context, message string, fs *Fields) {
	fs.fields = append(fs.fields, appNameField)
	_, msg, fields := executeMiddleLayers(ctx, message, fs)
	zapLogger.Info(msg, fields.fields...)
}

func Error(ctx context.Context, format string, a ...any) {
	fmt.Print(redColor)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Error(msg, fields.fields...)
	fmt.Print(defaultStyle)
}

func Warn(ctx context.Context, format string, a ...any) {
	fmt.Print(yellowColor)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Warn(msg, fields.fields...)
	fmt.Print(defaultStyle)
}

func Debug(ctx context.Context, format string, a ...any) {
	fmt.Print(greenColor)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Debug(msg, fields.fields...)
	fmt.Print(defaultStyle)
}

func Panic(ctx context.Context, format string, a ...any) {
	fmt.Print(redColor)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Panic(msg, fields.fields...)
	fmt.Print(defaultStyle)
}

func Fatal(ctx context.Context, format string, a ...any) {
	fmt.Print(redColor)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Fatal(msg, fields.fields...)
	fmt.Print(defaultStyle)
}
