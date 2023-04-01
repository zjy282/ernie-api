package ernieapi

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"testing"
)

const (
	fileContents = "This is a test file."
	boundary     = `MyBoundary`
)
const message = `
--MyBoundary
Content-Disposition: form-data; name="image"; filename="image.png"
Content-Type: text/plain

` + fileContents + `

--MyBoundary--
`

func TestClient_CreateTxt2Img(t *testing.T) {
	b := strings.NewReader(strings.ReplaceAll(message, "\n", "\r\n"))
	r := multipart.NewReader(b, boundary)
	f, err := r.ReadForm(0)
	if err != nil {
		t.Fatal("ReadForm:", err)
	}
	defer func() {
		_ = f.RemoveAll()
	}()

	client := NewClient("")
	ctx := context.Background()
	req := &Txt2ImgRequest{
		Text:       "睡莲",
		Style:      StyleOilPainting,
		Resolution: ResolutionSquareChart,
		Num:        1,
		Image:      f.File["image"][0],
	}

	response, err := client.CreateTxt2Img(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
