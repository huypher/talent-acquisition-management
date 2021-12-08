package user

import (
	"github.com/huypher/kit/container"

	"github.com/gin-gonic/gin"
	"github.com/huypher/kit/http_response"
	"github.com/huypher/talent-acquisition-management/auth"
)

func (d *userDelivery) getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := auth.UIDFromContext(c)
		if uid == (auth.UID{}) {
			http_response.BadRequest(c, "invalid request", nil)
			return
		}

		http_response.Success(c, "success", container.Map{
			"id":       uid.ID,
			"username": uid.Username,
			"name":     uid.Username,
		})
	}
}
