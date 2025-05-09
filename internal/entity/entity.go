package entity

type MessageRequest struct {
	Role        string   `json:"role"`
	Content     string   `json:"content"`
}

type MessageResponse struct {
	Content        string     `json:"content"`
}

type ChatCompletionRequest struct {
	Model          string           `json:"model"`
	Messages       []MessageRequest `json:"messages"`
}

type ChatCompletionResponse struct {
	Choices        []MessageResponse `json:"choices"`
}