package api

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
)

// Servers

func (c *Client) ListServers() (*ListServersResp, error) {
	body, err := c.GetApiResource("server")
	if err != nil {
		return nil, err
	}
	logrus.Debugf("API resp: %s", string(body))

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
	logrus.Debugf("API resp: %s", string(body))

	// json parsing
	var result GetServerResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// govalidator
	isValid, err := govalidator.ValidateStruct(result)
	if err != nil {
		return nil, err
	}
	if !isValid {
		logrus.Warnf("Structure seems invalid. Please report bug.")
	}

	return &result, nil
}

// Abuses

func (c *Client) ListAbuses() (*ListAbusesResp, error) {
	body, err := c.GetApiResource("abuse")
	if err != nil {
		return nil, err
	}
	logrus.Debugf("API resp: %s", string(body))

	// json parsing
	var result ListAbusesResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// govalidator
	isValid, err := govalidator.ValidateStruct(result)
	if err != nil {
		return nil, err
	}
	if !isValid {
		logrus.Warnf("Structure seems invalid. Please report bug.")
	}

	return &result, nil
}

// Users

func (c *Client) GetCurrentUser() (*GetCurrentUserResp, error) {
	body, err := c.GetApiResource("user")
	if err != nil {
		return nil, err
	}
	logrus.Debugf("API resp: %s", string(body))

	// json parsing
	var result GetCurrentUserResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// govalidator
	isValid, err := govalidator.ValidateStruct(result)
	if err != nil {
		return nil, err
	}
	if !isValid {
		logrus.Warnf("Structure seems invalid. Please report bug.")
	}

	return &result, nil
}
