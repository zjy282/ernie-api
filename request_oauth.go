package ernieapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const DefaultGrantType = "client_credentials"

type OAuthTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type OAuthTokenResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	AccessToken string `json:"data"`
}

type BCEOAuthTokenResponse struct {
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func CreateOAuthToken(ctx context.Context, request *OAuthTokenRequest) (response *OAuthTokenResponse, err error) {
	client := NewClientWithConfig(DefaultConfig(""))
	urlSuffix := "/oauth/token"

	if request.GrantType == "" {
		request.GrantType = DefaultGrantType
	}
	QueryParams := url.Values{}
	QueryParams.Add("grant_type", request.GrantType)
	QueryParams.Add("client_id", request.ClientID)
	QueryParams.Add("client_secret", request.ClientSecret)
	requestUrl := client.config.BaseURL + urlSuffix + "?" + QueryParams.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return
	}
	errResponse := &ResponseError{}
	err = client.sendRequest(req, &response, errResponse)
	return
}

func CreateBCEOAuthToken(ctx context.Context, request *OAuthTokenRequest) (response *BCEOAuthTokenResponse, err error) {
	client := NewClientWithConfig(DefaultBCEConfig(""))
	urlSuffix := "/oauth/2.0/token"

	if request.GrantType == "" {
		request.GrantType = DefaultGrantType
	}
	QueryParams := url.Values{}
	QueryParams.Add("grant_type", request.GrantType)
	QueryParams.Add("client_id", request.ClientID)
	QueryParams.Add("client_secret", request.ClientSecret)
	requestUrl := client.config.BaseURL + urlSuffix + "?" + QueryParams.Encode()
	fmt.Println(requestUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return
	}
	type BCEResponseError struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}

	errResponse := &BCEResponseError{}
	err = client.sendRequest(req, &response, errResponse)
	return
}
