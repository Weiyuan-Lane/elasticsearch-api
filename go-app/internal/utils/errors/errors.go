package errors

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var appName = os.Getenv("APP_NAME")

var (
	ErrPathNotFound               = initWrappedError(1, http.StatusNotFound)
	ErrUtilHttpRequestNewRequest  = initWrappedError(2, http.StatusInternalServerError)
	ErrUtilHttpRequestDo          = initWrappedError(3, http.StatusInternalServerError)
	ErrUtilHttpRequestStatusCode  = initWrappedError(4, http.StatusInternalServerError)
	ErrUtilHttpRequestDecoder     = initWrappedError(5, http.StatusInternalServerError)
	ErrCreateIndexNotAcknowledged = initWrappedError(6, http.StatusInternalServerError)
	ErrCreateIndexInvalidID       = initWrappedError(7, http.StatusInternalServerError)
	ErrIndexExistsInvalidStatus   = initWrappedError(8, http.StatusInternalServerError)
	ErrIndexNotFound              = initWrappedError(9, http.StatusNotFound)
	ErrIndexAlreadyCreated        = initWrappedError(10, http.StatusBadRequest)
)

type WrappedError struct {
	err        error
	statusCode int
}

func initWrappedError(errID int, httpStatus int) WrappedError {
	return WrappedError{
		err:        fmt.Errorf("%s.%d", appName, errID),
		statusCode: httpStatus,
	}
}

func (w WrappedError) Error() string {
	return w.err.Error()
}

func (w WrappedError) StatusCode() int {
	return w.statusCode
}
