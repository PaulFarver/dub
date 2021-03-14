package behindthename

import (
	"context"
	"encoding/json"
)

// RelatedNames
// This will return potential aliases for a given name.
// https://www.behindthename.com/api/help.php
func (c *ClientImpl) RelatedNames(ctx context.Context, name string) ([]byte, error) {
	req := c.r.R()
	req.SetContext(ctx)
	resp, err := req.Get("/related.json")
	if err != nil {
		return nil, err
	}

	return json.Marshal(resp.Result())
}
