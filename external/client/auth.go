package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rombintu/golearn/config"
	"github.com/rombintu/golearn/store"
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
	var user store.User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", err
	}
	return user.Password, nil
}
