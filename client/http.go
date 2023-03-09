package txclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/v7/httpclient"
)

// newHTTPClient creates a new heimdall HTTP client with retries and exponential backoff
func newHTTPClient(config Config) *httpclient.Client {
	backoff := heimdall.NewExponentialBackoff(2*time.Millisecond, 9*time.Millisecond, 2, 2*time.Millisecond)
	retrier := heimdall.NewRetrier(backoff)

	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(config.Timeout),
		httpclient.WithRetryCount(config.Retries),
		httpclient.WithRetrier(retrier),
	)
}

func performHttpReq[T any](method string, url string, header http.Header, body []byte, client *httpclient.Client) (T, error) {
	var data T

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}
	req.Header = header
	defer req.Body.Close()

	res, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, parseErrorResponse(res)
	}

	return data, json.NewDecoder(res.Body).Decode(&data)
}

// Parses the error reponse and returns the status code and error message
func parseErrorResponse(errResponse *http.Response) error {
	// default errMsg
	errMsg := fmt.Errorf("error: [%d] %s", errResponse.StatusCode, http.StatusText(errResponse.StatusCode))

	body, err := io.ReadAll(errResponse.Body)
	if err != nil {
		return errMsg
	}

	var errorMap map[string]string
	if err = json.Unmarshal(body, &errorMap); err != nil {
		return errMsg
	}

	if errorMessage, ok := errorMap["error"]; ok {
		errMsg = fmt.Errorf("%s: %s", errMsg, errorMessage)
	}

	return errMsg
}
