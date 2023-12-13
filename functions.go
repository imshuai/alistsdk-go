package alistsdk

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

const (
	DEFAULT_USERAGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"
	DEFAULT_TIMEOUT   = 30
)

var ()

func do(method string, url string, header map[string][]string, body io.Reader, timeout int, inscure bool) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(func() int {
			if timeout > 0 {
				return timeout
			} else {
				return DEFAULT_TIMEOUT
			}
		}()) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: inscure},
		},
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header = header
	request.Header.Set("User-Agent", DEFAULT_USERAGENT)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
