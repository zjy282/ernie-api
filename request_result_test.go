package ernie_api

import (
	"context"
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	client := NewClient("")
	ctx := context.Background()
	req := &ResultRequest{
		TaskId: 1,
	}

	response, err := client.GetResult(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
