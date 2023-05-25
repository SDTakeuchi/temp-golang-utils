package http

import (
	"fmt"
	"net/http"
	"net/url"
)

type (
	Endpoint interface {
		URL() *url.URL
		Method() string
	}
	endpoint struct {
		url    *url.URL
		method string
	}
)

var HTTPMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

func (e *endpoint) URL() *url.URL {
	return e.url
}

func (e *endpoint) Method() string {
	return e.method
}

func NewEndpoints(baseURL, path, method string) (Endpoint, error) {
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, err
	}
	if !isValidMethod(method) {
		return nil, fmt.Errorf("invalid method: %s", method)
	}
	return &endpoint{url: u, method: method}, nil
}

func isValidMethod(method string) bool {
	for _, m := range HTTPMethods {
		if method == m {
			return true
		}
	}
	return false
}
