package repository

import (
	"encoding/json"
	"fmt"

	"github.com/mackelaophu/ChatApp/base-service/internal/config"
	"github.com/mackelaophu/ChatApp/base-service/internal/model"
	"github.com/redis/go-redis/v9"
)

type ChatRepo struct {
}

func NewChatRepo() *ChatRepo {
	return &ChatRepo{}
}

// var users []model.User

func (Repo *ChatRepo) SaveUser(user *model.User) {
	key := "all_user"
	r := config.ConnectRedis()
	client := r.Rdb
	ctx := r.Ctx

	val, err := client.Get(*ctx, key).Result()
	var users []model.User

	if err == redis.Nil {
		// Lần đầu đăng ký
		users = append(users, *user)
	} else if err != nil {
		fmt.Println("Redis GET error:", err)
		return
	} else {
		// Key đã tồn tại → parse JSON
		if errParse := json.Unmarshal([]byte(val), &users); errParse != nil {
			fmt.Println("Invalid JSON:", errParse)
			return
		}
		// Kiểm tra trùng
		if contain(users, user) {
			fmt.Println("user already registered:", user.Id)
			return
		}
		users = append(users, *user)
		fmt.Println("register user:", user)
	}

	// Ghi lại vào Redis
	bytes, _ := json.Marshal(users)
	if err := client.Set(*ctx, key, bytes, 0).Err(); err != nil {
		fmt.Println("Redis SET error:", err)
	}
}

func (Repo *ChatRepo) GetAllUser() []model.User {
	key := "all_user" //config.AllUser
	r := config.ConnectRedis()
	client := r.Rdb
	ctx := r.Ctx
	val, err := client.Get(*ctx, key).Result()
	var users []model.User
	if err == redis.Nil {
		fmt.Println("no data")
		// Không có dữ liệu
		return []model.User{}
	} else if err != nil {
		fmt.Println("Redis GET error:", err)
		return []model.User{}
	}

	parseErr := json.Unmarshal([]byte(val), &users)
	if parseErr != nil {
		fmt.Println("Parse error:", parseErr)
		return []model.User{}
	}
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
