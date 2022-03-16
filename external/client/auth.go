package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rombintu/golearn/config"
)

type Client struct {
	Config *config.ConfigClient
}

func NewClient(config *config.ConfigClient) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) PingServer() (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/ping", c.Config.Default.Host, c.Config.Default.Port))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Client) GetToken(login, pass, role string) (string, error) {
	data := map[string]string{
		"account":  login,
		"password": pass,
	}

	if role != "" {
		data["role"] = role
	}

	json_data, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("http://%s:%s/auth", c.Config.Default.Host, c.Config.Default.Port)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return "", err
	}
	var res map[string]string

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res["token"], nil
}

func (c *Client) GetMyID() (string, error) {
	token := c.Config.Private.Token

	uri := fmt.Sprintf("http://%s:%s/user/token", c.Config.Default.Host, c.Config.Default.Port)

	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	resp.Header.Add("token", token)
	var res map[string]string

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res["token"], nil
}
