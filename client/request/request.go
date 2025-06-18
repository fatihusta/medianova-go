package request

import (
	"net/http"
	"net/url"
	"time"
)

const (
	CloudAPI    = "https://cloud.medianova.com"
	BaseAPIPath = "/api/v1"
)

type RequestConfig struct {
	httpClient     *http.Client
	BaseURL        *url.URL
	RequestTimeout time.Duration `yaml:"request_timeout,omitempty"`
}

func NewRequestConfig() *RequestConfig {
	base_url, _ := url.Parse(CloudAPI + BaseAPIPath)
	return &RequestConfig{
		BaseURL:        base_url,
		RequestTimeout: 5 * time.Second,
		httpClient:     &http.Client{},
	}
}

func (r *RequestConfig) GetClient() *http.Client {
	return r.httpClient
}

func (r *RequestConfig) SetClient(client *http.Client) {
	r.httpClient = client
}

func (r *RequestConfig) Do(req *http.Request) (*http.Response, error) {
	return r.httpClient.Do(req)
}
