package ernieapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	StyleAncient       = "古风"
	StyleAnime         = "二次元"
	StyleRealism       = "写实风格"
	StyleUkiyoE        = "浮世绘"
	StyleLowPoly       = "low poly"
	StyleFuturism      = "未来主义"
	StylePixel         = "像素风格"
	StyleConceptualArt = "概念艺术"
	StyleCyberpunk     = "赛博朋克"
	StyleLolita        = "洛丽塔风格"
	StyleBaroque       = "巴洛克风格"
	StyleSurrealism    = "超现实主义"
	StyleWatercolor    = "水彩画"
	StyleSteamPunk     = "蒸汽波艺术"
	StyleOilPainting   = "油画"
	StyleCartoon       = "卡通画"
)

const (
	ResolutionSquareChart     = "1024*1024"
	ResolutionLongChart       = "1024*1536"
	ResolutionHorizontalChart = "1536*1024"
)

type Txt2ImgRequest struct {
	Text       string   `json:"text" url:"text"`
	Style      string   `json:"style" url:"style"`
	Resolution string   `json:"resolution" url:"resolution"`
	Num        int      `json:"num" url:"num"`
	Image      *os.File `json:"image,omitempty" url:"image,omitempty"`
}

type Txt2ImgResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data Txt2ImgData `json:"data"`
}

type Txt2ImgData struct {
	TaskID    int    `json:"taskId"`
	RequestID string `json:"requestId"`
}

func (c *Client) CreateTxt2Img(ctx context.Context, request *Txt2ImgRequest) (response *Txt2ImgResponse, err error) {

	urlSuffix := "/rest/1.0/ernievilg/v1/txt2img"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()
	_ = writer.WriteField("text", request.Text)
	_ = writer.WriteField("style", request.Style)
	_ = writer.WriteField("resolution", request.Resolution)
	_ = writer.WriteField("num", fmt.Sprintf("%d", request.Num))

	if request.Image != nil {
		var dst io.Writer
		dst, err = writer.CreateFormFile("image", request.Image.Name())
		if err != nil {
			return
		}
		_, err = io.Copy(dst, request.Image)
		if err != nil {
			return
		}
	}

	fmt.Println(c.fullURL(urlSuffix))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.fullURL(urlSuffix), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "multipart/form-data")

	err = c.sendRequest(req, &response)
	return
}
