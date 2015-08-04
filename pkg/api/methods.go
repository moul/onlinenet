package api

import (
	"encoding/json"

	"github.com/docker/machine/log"
)

func (c *Client) ListServers() (*ListServersResp, error) {
	body, err := c.GetApiResource("server")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var result ListServersResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
