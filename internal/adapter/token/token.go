package token

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

var _ Token = (*token)(nil)

type Token interface {
	GetCurrentToken() string
}

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (t *token) GetCurrentToken() string {
	expiresAt := t.ExpiresAt
	apochNow := time.Now().Unix()
	timeDelta := apochNow - (expiresAt / 1000)
	if timeDelta > 0 {
		Auth()
	}
	token := t.AccessToken
	return token
}

func Auth() (*token, error) {
	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	method := "POST"
	payload := strings.NewReader("scope=GIGACHAT_API_PERS")
	uuid, _ := uuid.NewV4()
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	authKey, exists := os.LookupEnv("GPT_AUTHORIZATION_KEY")
	if !exists {
		return nil, errors.New("authorization key not found")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", uuid.String())
	req.Header.Add("Authorization", "Basic "+authKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var t *token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, ErrInvalidToken
	}
	return t, nil
}
