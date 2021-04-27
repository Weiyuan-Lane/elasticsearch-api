package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/logger"
)

type ApplicationConfig struct {
	Logger logger.Logger
}

func New() ApplicationConfig {
	return ApplicationConfig{
		Logger: logger.New(),
	}
}
