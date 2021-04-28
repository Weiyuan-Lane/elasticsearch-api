package httprequest

import (
	"bytes"
	"encoding/json"
	"net/http"

	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
)

const (
	getMethodName  string = "GET"
	postMethodName string = "POST"
	putMethodName  string = "PUT"
	headMethodName string = "HEAD"
)

func GetJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	body interface{},
	target interface{},
	errorTarget interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		getMethodName,
		queryParams,
		headers,
		body,
		target,
		errorTarget,
	)

	return statusCode, err
}

func PostJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	body interface{},
	target interface{},
	errorTarget interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		postMethodName,
		queryParams,
		headers,
		body,
		target,
		errorTarget,
	)

	return statusCode, err
}

func PutJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	body interface{},
	target interface{},
	errorTarget interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		putMethodName,
		queryParams,
		headers,
		body,
		target,
		errorTarget,
	)

	return statusCode, err
}

func HeadJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	body interface{},
	target interface{},
	errorTarget interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		headMethodName,
		queryParams,
		headers,
		body,
		target,
		errorTarget,
	)

	return statusCode, err
}

func invokeJSONRequest(
	url string,
	method string,
	queryParams map[string]string,
	headers map[string]string,
	body interface{},
	target interface{},
	errorTarget interface{},
) (int, error) {

	client := &http.Client{}
	var req *http.Request
	var err error

	var bodyBuffer *bytes.Buffer = nil
	if body != nil {
		byteArr, err := json.Marshal(body)
		if err != nil {
			return -1, apperrors.ErrUtilHttpRequestBodyMarshal
		}

		bodyBuffer = bytes.NewBuffer(byteArr)
		req, err = http.NewRequest(method, url, bodyBuffer)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return -1, apperrors.ErrUtilHttpRequestNewRequest
	}

	reqQuery := req.URL.Query()
	for queryKey, queryVal := range queryParams {
		reqQuery.Add(queryKey, queryVal)
	}
	req.URL.RawQuery = reqQuery.Encode()

	req.Header.Set("Content-Type", "application/json")
	for headerKey, headerVal := range headers {
		req.Header.Set(headerKey, headerVal)
	}

	res, err := client.Do(req)
	if err != nil {
		return -1, apperrors.ErrUtilHttpRequestDo
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode >= 400 && statusCode <= 599 {
		if errorTarget != nil {
			_ = json.NewDecoder(res.Body).Decode(errorTarget)
		}
		return statusCode, apperrors.ErrUtilHttpRequestStatusCode
	}

	if target != nil {
		err = json.NewDecoder(res.Body).Decode(target)
		if err != nil {
			return statusCode, apperrors.ErrUtilHttpRequestDecoder
		}
	}

	return statusCode, nil
}
