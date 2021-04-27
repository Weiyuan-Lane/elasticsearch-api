package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debugw(msg string, inputFields ...interface{})
	Errorw(msg string, inputFields ...interface{})
	Fatalw(msg string, inputFields ...interface{})
	Infow(msg string, inputFields ...interface{})
}

func New() Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()

	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
