package behindthename

import (
	"context"

	"github.com/pkg/errors"
)

type RelatedNamesParameters struct {
	Gender string
	Usage  string
}

type RelatedNamesResponse struct {
	ErrorCode int      `json:"error_code,omitempty" yaml:"error_code,omitempty"`
	Error     string   `json:"error,omitempty" yaml:"error,omitempty"`
	Names     []string `json:"names,omitempty" yaml:"names,omitempty"`
}

func (r *RelatedNamesResponse) GetNames() []string {
	return r.Names
}

// RelatedNames will return potential aliases for a given name.
// https://www.behindthename.com/api/help.php
func (c *ClientImpl) RelatedNames(ctx context.Context, name string, params RelatedNamesParameters) (*RelatedNamesResponse, error) {
	request := c.r.R()
	request.SetContext(ctx)
	request.SetResult(&RelatedNamesResponse{})

	request.SetQueryParam("name", name)
	if params.Gender != "" {
		request.SetQueryParam("gender", params.Gender)
	}
	if params.Usage != "" {
		request.SetQueryParam("usage", params.Usage)
	}

	resp, err := request.Get("/related.json")
	if err != nil {
		return nil, err
	}

	response := resp.Result().(*RelatedNamesResponse)

	if response.ErrorCode != 0 {
		return nil, errors.Errorf("Error response from behindthename: code=%v error=%s", response.ErrorCode, response.Error)
	}

	return response, nil
}
