package ernieapi

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
	"strings"
)

type Txt2ImgResultRequest struct {
	TaskId int `json:"taskId" url:"taskId"`
}

type Txt2ImgResultResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data Txt2ImgResultData `json:"data"`
}

type Txt2ImgResultData struct {
	Img        string                    `json:"img"`
	Waiting    string                    `json:"waiting"`
	ImgUrls    []Txt2ImgResultDataImages `json:"imgUrls"`
	CreateTime string                    `json:"createTime"`
	RequestID  string                    `json:"requestId"`
	Style      string                    `json:"style"`
	Text       string                    `json:"text"`
	Resolution string                    `json:"resolution"`
	TaskID     int                       `json:"taskId"`
	Status     int                       `json:"status"`
}

type Txt2ImgResultDataImages struct {
	Image string      `json:"image"`
	Score interface{} `json:"score"`
}

func (c *Client) GetTxt2ImgResult(ctx context.Context, request *Txt2ImgResultRequest) (response *Txt2ImgResultResponse, err error) {

	urlSuffix := "/rest/1.0/ernievilg/v1/getImg"

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
