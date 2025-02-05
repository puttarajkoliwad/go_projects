package logger

import "go.uber.org/zap"

var Log *zap.Logger

func init() {
	var err error
	if Log, err = zap.NewProduction(); err != nil {
		panic(err)
	}
}
