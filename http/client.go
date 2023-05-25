package http

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var defaultTimeout = time.Second * 10

type (
	Client interface {
		Do(e Endpoint, header, body map[string]string) (string, error)
	}
	client struct {
		timeout  time.Duration
	}
)

func NewClient(timeout time.Duration) Client {
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return &client{
		timeout:  timeout,
	}
}

func (c *client) Do(e Endpoint, header, body map[string]string) (string, error) {
	// init http client
	client := &http.Client{
		Timeout: c.timeout,
	}

	// prepare body
	var bodyString string
	for k, v := range body {
		bodyString += k + "=" + v + "&"
	}
	bodyString = strings.TrimSuffix(bodyString, "&")
	bodyParam := strings.NewReader(bodyString)

	req, err := http.NewRequest(
		e.Method(),
		e.URL().String(),
		bodyParam,
	)
	if err != nil {
		return "", err
	}

	// set header
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// read response
	defer resp.Body.Close()
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyResp), nil
}
