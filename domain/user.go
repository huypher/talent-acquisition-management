package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserUsecase interface {
	GetByUserName(ctx context.Context, username string) (User, error)
	GetByID(ctx context.Context, id int) (User, error)
}

type UserRepository interface {
	GetByUserName(ctx context.Context, username string) (User, error)
	GetByID(ctx context.Context, id int) (User, error)
}

type UserDelivery interface {
	Handler(c *gin.RouterGroup)
}
