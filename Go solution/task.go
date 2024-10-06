package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type GameResponse struct {
	Entry         string        `json:"entry"`
	EntryEmbedded EntryEmbedded `json:"entryEmbedded"`
}

type EntryEmbedded struct {
	EntryParams string `json:"entryParams"`
}

type RequestBody struct {
	UUID   string       `json:"uuid"`
	Player PlayerInfo   `json:"player"`
	Config ConfigParams `json:"config"`
}

type PlayerInfo struct {
	Id       string        `json:"id"`
	Update   bool          `json:"update"`
	Nickname string        `json:"nickname"`
	Language string        `json:"language"`
	Currency string        `json:"currency"`
	Session  SessionParams `json:"session"`
}

type SessionParams struct {
	Id string `json:"id"`
	Ip string `json:"ip"`
}

type ConfigParams struct {
	Brand BrandParams `json:"brand"`
}

type BrandParams struct {
	ID   string `json:"id"`
	Skin string `json:"skin"`
}

func main() {
	resp, err := authenticateUser()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Game response: %s\n", resp)
}

func authenticateUser() (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	proxyURL, err := url.Parse("http://goproxy.u1s1.biz:16600")
	if err != nil {
		return "", fmt.Errorf("invalid proxy URL: %v", err)
	}

	client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	endpoint := "https://skylinev88871.uat1.evo-test.com/ua/v1/skylinev88871001/test123"

	reqBody := RequestBody{
		UUID: "unique request identifier",
		Player: PlayerInfo{
			Id:       "a1a2a3a4",
			Update:   true,
			Nickname: "nickname",
			Language: "en-GB",
			Currency: "EUR",
			Session: SessionParams{
				Id: "111ssss3333rrrrr45555",
				Ip: "109.75.37.87",
			},
		},
		Config: ConfigParams{
			Brand: BrandParams{
				ID:   "1",
				Skin: "1",
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return "", fmt.Errorf("server returned non-200 status: %s", resp.Status)
	}

	var gameResponse GameResponse
	if err := json.NewDecoder(resp.Body).Decode(&gameResponse); err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	fmt.Printf("Game URL: %s\n", gameResponse.Entry)
	fmt.Printf("Embedded Entry Params: %s\n", gameResponse.EntryEmbedded.EntryParams)

	return gameResponse.Entry, nil
}
