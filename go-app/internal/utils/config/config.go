package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/logger"
)

type ApplicationConfig struct {
	Logger logger.Logger

	HTTPPort                    int
	HTTPGracefulShutdownSeconds int

	ElasticsearchHost     string
	ElasticsearchPort     int
	ElasticsearchProtocol string
}

func New() ApplicationConfig {
	return ApplicationConfig{
		Logger:                      logger.New(),
		HTTPPort:                    envVarAsInt("HTTP_PORT"),
		HTTPGracefulShutdownSeconds: envVarAsInt("HTTP_GRACEFUL_SHUTDOWN_SECONDS"),
		ElasticsearchHost:           envVarAsStr("ELASTICSEARCH_HOST"),
		ElasticsearchPort:           envVarAsInt("ELASTICSEARCH_PORT"),
		ElasticsearchProtocol:       envVarAsStr("ELASTICSEARCH_PROTOCOL"),
	}
}

func envVarAsStr(envName string) string {
	valueStr := os.Getenv(envName)
	return valueStr
}

func envVarAsInt(envName string) int {
	valueStr := os.Getenv(envName)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(err)
	}

	return value
}

func envVarAsBool(envName string) bool {
	valueStr := os.Getenv(envName)
	if valueStr == "true" {
		return true
	} else if valueStr == "false" {
		return false
	}

	panic(fmt.Sprintf("\"%s\" env value is not boolean", envName))
}
