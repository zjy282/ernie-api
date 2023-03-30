package ernieapi

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateV3Customize(t *testing.T) {
	client := NewClient("")
	ctx := context.Background()
	req := &V3CustomizeRequest{
		Async:            1,
		Text:             "标题：芍药香氛的沐浴乳\\n文案：",
		MinDecLen:        32,
		SeqLen:           512,
		TopP:             0.9,
		TaskPrompt:       TaskPromptAdText,
		PenaltyScore:     1.2,
		IsUnidirectional: 0,
		TypeId:           1,
	}

	response, err := client.CreateV3Customize(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
