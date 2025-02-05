package logger

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	var err error
	if log, err = zap.NewProduction(); err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Fields) {
	log.Info(message, fields...)
}
