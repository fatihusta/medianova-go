package request

import (
	"net/http"
	"net/url"
	"time"
)

const (
	CloudAPIv1 = "https://cloud.medianova.com/api/v1"
	CloudAPIv2 = "https://cloud.medianova.com/api/v2"
)

type RequestConfig struct {
	httpClient     *http.Client
	BaseURL        *url.URL
	RequestTimeout time.Duration `yaml:"request_timeout,omitempty"`
}

func NewRequestConfig() *RequestConfig {
	base_url, _ := url.Parse(CloudAPIv1)
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
