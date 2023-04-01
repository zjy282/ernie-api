package ernieapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_GetTxt2ImgResult(t *testing.T) {
	client := NewClient("")
	ctx := context.Background()
	req := &Txt2ImgResultRequest{
		TaskId: 1,
	}

	response, err := client.GetTxt2ImgResult(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
