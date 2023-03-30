package ernie_api

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
)

type ResultRequest struct {
	TaskId int `json:"taskId" url:"taskId"`
}

type ResultResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ResultData `json:"data"`
}

type ResultData struct {
	Result     string `json:"result"`
	CreateTime string `json:"createTime"`
	RequestID  string `json:"requestId"`
	Text       string `json:"text"`
	TaskID     int    `json:"taskId"`
	Status     int    `json:"status"`
}

func (c *Client) GetResult(ctx context.Context, request *ResultRequest) (response *ResultResponse, err error) {

	urlSuffix := "/rest/1.0/ernie/v1/getResult"

	requestParams, err := query.Values(*request)
	if err != nil {
		return response, ErrV3CustomizeRequest
	}

	req, err := c.requestBuilder.build(ctx, http.MethodPost, c.fullURL(urlSuffix), requestParams)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
