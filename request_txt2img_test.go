package ernieapi

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_CreateTxt2Img(t *testing.T) {
	filePath := "./test.png"
	file, err := os.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	client := NewClient("")
	ctx := context.Background()
	req := &Txt2ImgRequest{
		Text:       "睡莲",
		Style:      StyleOilPainting,
		Resolution: ResolutionSquareChart,
		Num:        1,
		Image:      file,
	}

	response, err := client.CreateTxt2Img(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
