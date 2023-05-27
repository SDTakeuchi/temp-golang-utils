package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var defaultTimeout = time.Second * 10

type (
	Client interface {
		Do(e Endpoint, header, body map[string]string) (string, error)
	}
	client struct {
		timeout time.Duration
	}
)

func NewClient(timeout time.Duration) Client {
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return &client{
		timeout: timeout,
	}
}

// Do executes the request, sends body in json style
func (c *client) Do(e Endpoint, header, body map[string]string) (string, error) {
	// init http client
	client := &http.Client{
		Timeout: c.timeout,
	}

	// prepare body
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	fmt.Printf("request body: %s\n", string(bodyBytes))
	param := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(
		e.Method(),
		e.URL().String(),
		param,
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
