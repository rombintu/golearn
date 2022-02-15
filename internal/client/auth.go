package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PingServer(serverUrl string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/ping", serverUrl))
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

func GetToken(serverUrl, account, pass, role string) (string, error) {
	data := map[string]string{
		"account":  account,
		"password": pass,
	}

	if role != "" {
		data["role"] = role
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("http://%s/auth", serverUrl)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return "", err
	}
	var res map[string]string

	json.NewDecoder(resp.Body).Decode(&res)

	return res["token"], nil
}
