package ernieapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	config ClientConfig

	requestBuilder requestBuilder
}

func NewClient(accessToken string) *Client {
	config := DefaultConfig(accessToken)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: newRequestBuilder(),
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ResponseError
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil {
			reqErr := RequestError{
				StatusCode: res.StatusCode,
				Err:        err,
			}
			return fmt.Errorf("error, %w", &reqErr)
		}
		return fmt.Errorf("error, http code: %d, message: %w", res.StatusCode, &errRes)
	}

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s?access_token=%s", c.config.BaseURL, suffix, c.config.accessToken)
}
