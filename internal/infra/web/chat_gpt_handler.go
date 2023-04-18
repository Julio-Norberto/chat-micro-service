package web

import "github.com/Julio-Norberto/chat-micro-service/internal/use-case/chatcompletionstream"

type WebChatGPTHandler struct {
	CompletionUseCase chatcompletionstream.ChatCompletionUseCase
}
