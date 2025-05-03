package repository

import (
	"fmt"

	"github.com/mackelaophu/ChatApp/base-service/internal/model"
)

type ChatRepo struct {
}

func NewChatRepo() *ChatRepo {
	return &ChatRepo{}
}

var users []model.User

func (Repo *ChatRepo) SaveUser(user *model.User) {
	if !contain(users, user) {
		fmt.Println("register user:", user)
		users = append(users, *user)
	} else {
		fmt.Println("user already registered:", user.Id)
	}
}

func (Repo *ChatRepo) GetAllUser() []model.User {
	return users
}

func contain(users []model.User, user *model.User) bool {
	for _, v := range users {
		if v.Id == user.Id {
			return true
		}
	}
	return false
}
