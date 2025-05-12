package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"crypto/tls"

	"gpt-service/internal/adapter/token"
	"gpt-service/internal/entity"
)

var (
	ErrSendRequest   = fmt.Errorf("failed to send request")
	ErrCreateRequest = fmt.Errorf("failed to create request")
)

var _ Handler = (*handler)(nil)

type Handler interface {
	GetGPTRecommendation([]string, string) (string, error)
	GetGPTImage(string) ([]byte, string, error)
}

type handler struct {
	accessToken token.Token
}

// Конструктор для создания нового обработчика
func New(token token.Token) Handler {
	return &handler{
		accessToken: token,
	}
}

func (h *handler) GetGPTRecommendation(products []string, preference string) (string, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url := "https://gigachat.devices.sberbank.ru//api/v1/chat/completions"
	method := "POST"
	prompt := fmt.Sprintf("У меня есть следующие продукты: %s и следующее предпочтение по калорийности: %s Что я могу из них приготовить?", strings.Join(products, ","), preference)
	requestBody := entity.ChatCompletionRequest{
		Model: "GigaChat",
		Messages: []entity.MessageRequest{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+h.accessToken.GetCurrentToken())
	if err != nil {
		return "", ErrCreateRequest
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", ErrSendRequest
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка API: %s, тело ответа: %s", res.Status, string(body))
	}

	return string(body), nil
}

// TODO: заглушка, нужно реализовать получение изображения
func (h *handler) GetGPTImage(message string) ([]byte, string, error) {
	return nil, "", nil
}
