package service

import (
	"github.com/mackelaophu/ChatApp/base-service/internal/model"
	"github.com/mackelaophu/ChatApp/base-service/internal/repository"
)

type ChatService struct {
	Repo *repository.ChatRepo
}

func NewChatService(Repo *repository.ChatRepo) *ChatService {
	return &ChatService{Repo: Repo}
}

func (Service *ChatService) SaveUser(user *model.User) {
	Service.Repo.SaveUser(user)
}

func (Service *ChatService) GetAllUser() []model.User {
	return Service.Repo.GetAllUser()
}
