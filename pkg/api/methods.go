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
	if err = json.Unmarshal(body, &result); err != nil {
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
	if err = json.Unmarshal(body, &result); err != nil {
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

func (c *Client) RebootServer(identifier int, reason, email string) (*RebootServerResp, error) {
	payload := map[string]string{
		"reason": reason,
		"email":  email,
	}
	body, err := c.PostApiResource(fmt.Sprintf("server/reboot/%d", identifier), payload)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("API resp: %s", string(body))

	// json parsing
	var result RebootServerResp
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// govalidator
	// FIXME: to do

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
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// govalidator
	/*
		isValid, err := govalidator.ValidateStruct(result)
		if err != nil {
			return nil, err
		}
		if !isValid {
			logrus.Warnf("Structure seems invalid. Please report bug.")
		}
	*/

	return &result, nil
}

// Ddos

func (c *Client) ListDdos() (*ListDdosResp, error) {
	body, err := c.GetApiResource("network/ddos")
	if err != nil {
		return nil, err
	}
	logrus.Debugf("API resp: %s", string(body))

	// json parsing
	var result ListDdosResp
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// govalidator
	/*
		isValid, err := govalidator.ValidateStruct(result)
		if err != nil {
			return nil, err
		}
		if !isValid {
			logrus.Warnf("Structure seems invalid. Please report bug.")
		}
	*/

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
	if err = json.Unmarshal(body, &result); err != nil {
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
