package ernieapi

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateOAuthToken(t *testing.T) {
	ctx := context.Background()
	req := &OAuthTokenRequest{
		ClientID:     "test",
		ClientSecret: "test",
	}
	response, err := CreateOAuthToken(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("response=%+v\n", response))
}
