package service

import "github.com/mackelaophu/ChatApp/base-service/internal/repository"

type ChatService struct {
	Repo *repository.ChatRepo
}

func NewChatService(Repo *repository.ChatRepo) *ChatService {
	return &ChatService{Repo: Repo}
}
