package api

import (
	"encoding/json"
	"fmt"

	"github.com/docker/machine/log"
)

// Servers

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

func (c *Client) GetServer(identifier int) (*GetServerResp, error) {
	body, err := c.GetApiResource(fmt.Sprintf("server/%d", identifier))
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var result GetServerResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Abuses

func (c *Client) ListAbuses() (*ListAbusesResp, error) {
	body, err := c.GetApiResource("abuse")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var result ListAbusesResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Users

func (c *Client) GetCurrentUser() (*GetCurrentUserResp, error) {
	body, err := c.GetApiResource("user")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var result GetCurrentUserResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
