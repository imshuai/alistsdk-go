package alistsdk

import (
	"io"
	"net/http"
	"time"
)

const (
	DEFAULT_USERAGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"
	DEFAULT_TIMEOUT   = 30
)

var ()

func post(url string, header map[string][]string, body []byte) ([]byte, error) {
	return postWithTimeout(url, header, body, DEFAULT_TIMEOUT)
}

func postWithTimeout(url string, header map[string][]string, body []byte, timeout int) ([]byte, error) {

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	request := &http.Request{}
	request.Header = header
	request.Method = "POST"
	request.Header.Set("User-Agent", DEFAULT_USERAGENT)
	request.URL, _ = request.URL.Parse(url)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func get(url string, header map[string][]string) ([]byte, error) {
	client := &http.Client{}
	request := &http.Request{}
	request.Header = header
	request.Method = "GET"
	request.Header.Set("User-Agent", DEFAULT_USERAGENT)
	request.URL, _ = request.URL.Parse(url)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
