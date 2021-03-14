package behindthename

import (
	"context"
	"encoding/json"
)

// RandomName
// This will return a random name.
// https://www.behindthename.com/api/help.php
func (c *ClientImpl) RandomName(ctx context.Context, name string) ([]byte, error) {
	req := c.r.R()
	req.SetContext(ctx)
	resp, err := req.Get("/random.json")
	if err != nil {
		return nil, err
	}

	return json.Marshal(resp.Result())
}
