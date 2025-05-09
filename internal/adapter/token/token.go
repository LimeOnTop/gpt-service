package token

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func New(AuthorizationKey string) (*Token, error) {
	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	method := "POST"
	payload := strings.NewReader("scope=GIGACHAT_API_PERS")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", uuid.New().String())
	req.Header.Add("Authorization", "Basic "+AuthorizationKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, ErrInvalidToken
	}
	return &t, nil
}
