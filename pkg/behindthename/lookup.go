package behindthename

import (
	"context"
	"encoding/json"
)

// Lookup
// This will return information about a given name.
// https://www.behindthename.com/api/help.php
func (c *ClientImpl) Lookup(ctx context.Context) ([]byte, error) {
	req := c.r.R()
	req.SetContext(ctx)
	resp, err := req.Get("/lookup.json")
	if err != nil {
		return nil, err
	}

	return json.Marshal(resp.Result())
}
