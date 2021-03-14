package behindthename

import (
	"context"
	"net/http"

	"gopkg.in/resty.v1"
)

// Client is an interface to behindthenames api
type Client interface {
	Lookup(ctx context.Context, name string) (result []byte, err error)
	RandomName(ctx context.Context) (result []byte, err error)
	RelatedNames(ctx context.Context, name string) (result []byte, err error)
}

// ClientImpl is an implementation of the poeditor client interface
type ClientImpl struct {
	r *resty.Client
}

// NewClient creates a new poeditor api client
func NewClient(apiToken string, httpClient *http.Client) *ClientImpl {
	r := resty.NewWithClient(httpClient)
	r.SetQueryParam("key", apiToken)

	r.SetHostURL("https://www.behindthename.com/api/")

	return &ClientImpl{
		r: r,
	}
}
