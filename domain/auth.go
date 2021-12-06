package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthUsecase interface {
	Login(ctx context.Context, username string, password string) (string, error)
}

type AuthDelivery interface {
	Handler(c *gin.RouterGroup)
}
