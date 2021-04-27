package httprequest

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
)

const (
	getMethodName  string = "GET"
	putMethodName  string = "PUT"
	headMethodName string = "HEAD"
)

func GetJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	target interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		getMethodName,
		queryParams,
		headers,
		target,
	)

	return statusCode, err
}

func PutJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	target interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		putMethodName,
		queryParams,
		headers,
		target,
	)

	return statusCode, err
}

func HeadJSON(
	url string,
	queryParams map[string]string,
	headers map[string]string,
	target interface{},
) (int, error) {

	statusCode, err := invokeJSONRequest(
		url,
		headMethodName,
		queryParams,
		headers,
		target,
	)

	return statusCode, err
}

func invokeJSONRequest(
	url string,
	method string,
	queryParams map[string]string,
	headers map[string]string,
	target interface{},
) (int, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return -1, apperrors.ErrUtilHttpRequestNewRequest
	}

	reqQuery := req.URL.Query()
	for queryKey, queryVal := range queryParams {
		reqQuery.Add(queryKey, queryVal)
	}
	req.URL.RawQuery = reqQuery.Encode()

	for headerKey, headerVal := range headers {
		req.Header.Set(headerKey, headerVal)
	}

	res, err := client.Do(req)
	if err != nil {
		return -1, apperrors.ErrUtilHttpRequestDo
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode >= 500 && statusCode <= 599 {
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
