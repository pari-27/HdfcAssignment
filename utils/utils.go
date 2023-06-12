package utils

import (
	"net/http"
	"net/url"
	"time"
)

func SendHTTPRequest(websiteName string, method string) (err error) {
	u, err := url.Parse(websiteName)
	if err != nil {
		return
	}
	url := url.URL{Scheme: "http", Host: u.Host, Path: u.Path}
	request, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	_, err = httpClient.Do(request)
	if err != nil {
		return
	}
	return
}
