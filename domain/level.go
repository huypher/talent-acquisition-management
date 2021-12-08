package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	tam "github.com/pghuy/talent-acquisition-management"
)

type Level struct {
	ID   int           `json:"id"`
	Code tam.LevelType `json:"code"`
	Name string        `json:"name"`
}

type LevelUsecase interface {
	GetAll(ctx context.Context) ([]Level, error)
}

type LevelRepository interface {
	GetAll(ctx context.Context) ([]Level, error)
}

type LevelDelivery interface {
	Handler(c *gin.RouterGroup)
}
