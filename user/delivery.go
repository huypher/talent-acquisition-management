package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/talent-acquisition-management/domain"
)

type userDelivery struct {
	userUsecase domain.UserUsecase
}

func NewProductDelivery(userUsecase domain.UserUsecase) *userDelivery {
	return &userDelivery{
		userUsecase: userUsecase,
	}
}

func (d *userDelivery) Handler(c *gin.RouterGroup) {
	c.Handle(http.MethodGet, "/user", d.getUser())
}
