package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/EgorMizerov/testMascotGaming/internal/client/domain"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	httpClient *http.Client
	api        string
}

func NewClient(api string) *Client {
	cert, err := tls.LoadX509KeyPair("./ssl/client.crt", "./ssl/client.key")
	if err != nil {
		log.Fatalf("error loadkeys: %e", err)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return &Client{httpClient: client, api: api}
}

func (c *Client) GetGameList() (domain.GameList, error) {
	query := `{
		"jsonrpc":"2.0",
		"method":"Game.List",
		"id":"11"
	}`

	req, err := http.NewRequest("POST", c.api, strings.NewReader(query))
	if err != nil {
		return domain.GameList{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return domain.GameList{}, err
	}
	var gameList domain.GameList
	json.NewDecoder(res.Body).Decode(&gameList)
	return gameList, err
}

func (c *Client) SetBankGroup(userId string) (string, error) {
	bankId := uuid.New().String()
	query := fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"Game.List",
		"id":"11",
		"params":{
			"Id":"%s",
			"Currency": "USD"
		}
	}`, bankId)

	req, err := http.NewRequest("POST", c.api, strings.NewReader(query))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	_, err = c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	return bankId, nil
}

func (c *Client) SetPlayer(userId, username, bankId string) error {
	query := fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"Player.Set",
		"id":11,
		"params":{
			"Id":"%s",
			"Nick":"%s",
			"BankGroupId":"%s"
		}
	}`, userId, username, bankId)

	req, err := http.NewRequest("POST", c.api, strings.NewReader(query))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	_, err = c.httpClient.Do(req)

	return err
}

func (c *Client) StartSession(playerId, gameId string) (string, string, error) {
	query := fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"Session.Create",
		"id":11,
		"params":{
			"PlayerId":"%s",
			"GameId":"%s"
		}
	}`, playerId, gameId)

	req, err := http.NewRequest("POST", c.api, strings.NewReader(query))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", "", err
	}

	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))

	var SSResp domain.StartSessionResponse
	err = json.Unmarshal(r, &SSResp)

	return SSResp.Result.SessionId, SSResp.Result.SessionUrl, err
}

func (c *Client) StartDemoSession(bankGroupId, gameId string) (string, string, error) {
	query := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "Session.CreateDemo",
		"id": 321864203,
		"params": {
			"BankGroupId": "%s",
			"GameId": "%s",
			"StartBalance": 10000
		}
	}`, bankGroupId, gameId)

	req, err := http.NewRequest("POST", c.api, strings.NewReader(query))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", "", err
	}

	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))

	var SSResp domain.StartSessionResponse
	err = json.Unmarshal(r, &SSResp)

	return SSResp.Result.SessionId, SSResp.Result.SessionUrl, err
}
