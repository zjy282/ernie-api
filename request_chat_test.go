package ernieapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_CreateChat(t *testing.T) {
	client := NewClientWithConfig(DefaultBCEConfig("test"))
	ctx := context.Background()
	req := &ChatRequest{
		User: "test",
		Messages: []ChatRequestMessage{
			{Role: ChatRoleUser, Content: "介绍一下你自己"},
		},
	}

	response, err := client.CreateChat(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
