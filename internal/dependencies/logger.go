package dependencies

import "go.uber.org/zap"

// Logger is a logger
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
}
