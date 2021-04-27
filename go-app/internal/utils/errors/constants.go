package errors

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var appName = os.Getenv("APP_NAME")

var (
	ErrPathNotFound = fmt.Errorf("%s.%d", appName, 1)
)
