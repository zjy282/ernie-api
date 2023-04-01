package ernieapi

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
	"strings"
)

type V3CustomizeResultRequest struct {
	TaskId int `json:"taskId" url:"taskId"`
}

type V3CustomizeResultResponse struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data V3CustomizeResultData `json:"data"`
}

type V3CustomizeResultData struct {
	Result     string `json:"result"`
	CreateTime string `json:"createTime"`
	RequestID  string `json:"requestId"`
	Text       string `json:"text"`
	TaskID     int    `json:"taskId"`
	Status     int    `json:"status"`
}

func (c *Client) GetV3CustomizeResult(ctx context.Context, request *V3CustomizeResultRequest) (response *V3CustomizeResultResponse, err error) {

	urlSuffix := "/rest/1.0/ernie/v1/getResult"

	requestParams, err := query.Values(*request)
	if err != nil {
		return response, ErrV3CustomizeRequest
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.fullURL(urlSuffix), strings.NewReader(requestParams.Encode()))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
