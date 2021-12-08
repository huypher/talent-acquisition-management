package level

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/talent-acquisition-management/domain"
)

type levelDelivery struct {
	userUsecase domain.LevelUsecase
}

func NewLevelDelivery(levelUsecase domain.LevelUsecase) *levelDelivery {
	return &levelDelivery{
		userUsecase: levelUsecase,
	}
}

func (d *levelDelivery) Handler(c *gin.RouterGroup) {
	c.Handle(http.MethodGet, "/levels", nil)
}
