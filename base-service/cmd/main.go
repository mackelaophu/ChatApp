package main

import (
	"github.com/mackelaophu/ChatApp/base-service/internal/handler"
	"github.com/mackelaophu/ChatApp/base-service/internal/repository"
	"github.com/mackelaophu/ChatApp/base-service/internal/service"
)

func main() {
	repo := repository.NewChatRepo()
	service := service.NewChatService(repo)
	handler := handler.NewChatHandler(service)
	handler.Init()
}
