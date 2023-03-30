package ernie_api

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type requestBuilder interface {
	build(ctx context.Context, method, url string, request url.Values) (*http.Request, error)
}

type httpRequestBuilder struct {
	marshaller marshaller
}

func newRequestBuilder() *httpRequestBuilder {
	return &httpRequestBuilder{
		marshaller: &jsonMarshaller{},
	}
}

func (b *httpRequestBuilder) build(ctx context.Context, method, url string, request url.Values) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	return http.NewRequestWithContext(
		ctx,
		method,
		url,
		strings.NewReader(request.Encode()),
	)
}
