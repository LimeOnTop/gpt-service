package handler

import (
	"io"
	"net/http"
	"strings"
)

var _ Handler = (*handler)(nil)

type Handler interface{
	GetGPTRecommendation([]string, string)(string, error)
	GetGPTImage(string)([]byte)
}

type handler struct {
	
}

//Конструктор для создания нового обработчика
func New() Handler {
	return &handler{}
}

func (h *handler) GetGPTRecommendation(products []string, preference string) (string, error) {
		
	return Message, nil
}

func (h *handler) GetGPTImage(message string) ([]byte, string, error) {

}

