package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger(logFilePath string) error {
	// Open the file to append logs
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// Configure the log file writer
	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(zapcore.Lock(os.Stdout)), // Log to console
		zapcore.AddSync(logFile),                 // Log to file
	)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Create the core
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	// Construct the Logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}
