package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var Logger *zap.SugaredLogger

var (
	mappingLogLevel = map[string]zapcore.Level{
		"DEBUG": zapcore.DebugLevel,
		"INFO":  zapcore.InfoLevel,
		"WARN":  zapcore.WarnLevel,
		"ERROR": zapcore.ErrorLevel,
	}
)

func SetUpLogger() {
	logLevel := getLogLevel()
	encoder := getEncoder()
	logWriter := getLogWriter()
	coreLogger := zapcore.NewCore(encoder, logWriter, logLevel)
	Logger = zap.New(coreLogger).Sugar()
}

func GetLogger() *zap.SugaredLogger{
	return Logger
}

func getLogLevel() zapcore.Level {
	var level string
	if value, ok := os.LookupEnv("LOGLEVEL"); ok {
		level = strings.ToUpper(value)
	} else {
		level = "INFO"
	}
	return mappingLogLevel[level]
}

func getEncoder() zapcore.Encoder {
	encoderConfig := getEncoderConfig()
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.ConsoleSeparator = " "
	return encoderConfig
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
