package auth

import (
	"errors"

	"github.com/huypher/kit/container"
	httpresponse "github.com/huypher/kit/http_response"

	"github.com/gin-gonic/gin"
)

func (d *authDelivery) loginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		err := c.BindJSON(&req)
		if err != nil {
			httpresponse.Error(c, err)
			return
		}

		token, err := d.authUsecase.Login(c, req.Username, req.Password)
		if err != nil {
			var invalidUserNameError *InvalidUserNameError
			if errors.As(err, &invalidUserNameError) {
				httpresponse.NotAuthorized(c, err.Error())
				return
			}

			var wrongPassError *WrongPasswordError
			if errors.As(err, &wrongPassError) {
				httpresponse.NotAuthorized(c, err.Error())
				return
			}

			httpresponse.Error(c, err)
			return
		}

		httpresponse.Success(c, "login success", container.Map{"token": token})
	}
}
