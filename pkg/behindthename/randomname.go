package behindthename

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
)

type RandomNameParameters struct {
	Gender        string
	Usage         string
	Number        int
	RandomSurname bool
}

type RandomNameResponse struct {
	ErrorCode int      `json:"error_code,omitempty" yaml:"error_code,omitempty"`
	Error     string   `json:"error,omitempty" yaml:"error,omitempty"`
	Names     []string `json:"names,omitempty" yaml:"names,omitempty"`
}

func (r *RandomNameResponse) GetNames() []string {
	return r.Names
}

// RandomName will return a random name.
// https://www.behindthename.com/api/help.php
func (c *ClientImpl) RandomName(ctx context.Context, params RandomNameParameters) (*RandomNameResponse, error) {
	request := c.r.R()
	request.SetContext(ctx)
	request.SetResult(&RandomNameResponse{})

	request.SetQueryParam("number", strconv.Itoa(params.Number))
	if params.Gender != "" {
		request.SetQueryParam("gender", params.Gender)
	}
	if params.Usage != "" {
		request.SetQueryParam("usage", params.Usage)
	}
	if params.RandomSurname {
		request.SetQueryParam("randomsurname", "yes")
	}

	resp, err := request.Get("/random.json")
	if err != nil {
		return nil, err
	}

	response := resp.Result().(*RandomNameResponse)

	if response.ErrorCode != 0 {
		return nil, errors.Errorf("Error response from behindthename: code=%v error=%s", response.ErrorCode, response.Error)
	}

	return response, nil
}
