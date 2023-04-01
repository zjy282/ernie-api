package ernieapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_GetV3CustomizeResult(t *testing.T) {
	client := NewClient("")
	ctx := context.Background()
	req := &V3CustomizeResultRequest{
		TaskId: 1,
	}

	response, err := client.GetV3CustomizeResult(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
