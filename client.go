package ernieapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	apiURLv1    = "https://wenxin.baidu.com/moduleApi/portal/api"
	apiURLv1BCE = "https://aip.baidubce.com"
)

type APIType string

const (
	APITypeWX  APIType = "wenxin"
	APITypeBCE APIType = "bce"
)

type ClientConfig struct {
	accessToken string
	HTTPClient  *http.Client
	BaseURL     string
	APIType     APIType
}

func DefaultConfig(accessToken string) ClientConfig {
	return ClientConfig{
		HTTPClient:  &http.Client{},
		BaseURL:     apiURLv1,
		accessToken: accessToken,
		APIType:     APITypeWX,
	}
}

func DefaultBCEConfig(accessToken string) ClientConfig {
	return ClientConfig{
		HTTPClient:  &http.Client{},
		BaseURL:     apiURLv1BCE,
		accessToken: accessToken,
		APIType:     APITypeBCE,
	}
}

type Client struct {
	config ClientConfig
}

func NewClient(accessToken string) *Client {
	config := DefaultConfig(accessToken)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) sendRequest(req *http.Request, successResponse, errResponse interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		err = json.NewDecoder(res.Body).Decode(errResponse)
		if err != nil {
			reqErr := RequestError{
				StatusCode: res.StatusCode,
				Err:        err,
			}
			return fmt.Errorf("error, %+v", reqErr)
		}
		return fmt.Errorf("error, http code: %d, message: %+v", res.StatusCode, errResponse)
	}

	if successResponse != nil {
		if err = json.NewDecoder(res.Body).Decode(successResponse); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s?access_token=%s", c.config.BaseURL, suffix, c.config.accessToken)
}

var (
	ErrV3CustomizeRequest = errors.New("request params convert error")
)

type RequestError struct {
	StatusCode int
	Err        error
}

type ResponseError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Err  error
}
