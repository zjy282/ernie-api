package ernie_api

import (
	"net/http"
)

const (
	apiURLv1 = "https://wenxin.baidu.com/moduleApi/portal/api"
)

type ClientConfig struct {
	accessToken string
	HTTPClient  *http.Client
	BaseURL     string
}

func DefaultConfig(accessToken string) ClientConfig {
	return ClientConfig{
		HTTPClient:  &http.Client{},
		BaseURL:     apiURLv1,
		accessToken: accessToken,
	}
}
