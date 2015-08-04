package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

type Client struct {
	Token  string
	client *http.Client
}

func NewClient() Client {
	return NewClientWithToken(os.Getenv("ONLINENET_TOKEN"))
}

func NewClientWithToken(token string) Client {
	return Client{
		Token:  token,
		client: &http.Client{},
	}
}

func (c *Client) GetApiResource(resource string) ([]byte, error) {
	url := fmt.Sprintf("https://api.online.net/api/v1/%s", resource)
	logrus.Debugf("url: %q", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
