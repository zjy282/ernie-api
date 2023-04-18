package ernieapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ChatRoleUser      = "user"
	ChatRoleAssistant = "assistant"
)

type ChatRequest struct {
	User     string               `json:"user"`
	Stream   bool                 `json:"stream"`
	Messages []ChatRequestMessage `json:"messages"`
}

type ChatRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	ID               string            `json:"id"`
	Object           string            `json:"object"`
	Created          int               `json:"created"`
	Result           string            `json:"result"`
	NeedClearHistory bool              `json:"need_clear_history"`
	Usage            ChatResponseUsage `json:"usage"`
}

type ChatResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (c *Client) CreateChat(ctx context.Context, request *ChatRequest) (response *ChatResponse, err error) {
	urlSuffix := "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions"

	request.Stream = false
	rawData, err := json.Marshal(request)
	if err != nil {
		return
	}
	fmt.Println(string(rawData))
	fmt.Println(c.fullURL(urlSuffix))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.fullURL(urlSuffix), bytes.NewBuffer(rawData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	type ResponseError struct {
		ErrorCode int    `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}
	errResponse := &ResponseError{}
	err = c.sendRequest(req, &response, errResponse)
	return
}
