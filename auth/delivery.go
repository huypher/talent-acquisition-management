package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huypher/talent-acquisition-management/domain"
)

const (
	loginPath    = "/login"
	userInfoPath = "/talent"
)

type authDelivery struct {
	authUsecase domain.AuthUsecase
}

func NewAuthDelivery(authUsecase domain.AuthUsecase) *authDelivery {
	return &authDelivery{
		authUsecase: authUsecase,
	}
}

func (d *authDelivery) Handler(c *gin.RouterGroup) {
	group := c.Group("auth")
	group.Handle(http.MethodPost, loginPath, d.loginHandler())
}
