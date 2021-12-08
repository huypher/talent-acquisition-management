package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/talent-acquisition-management/auth"
	"github.com/pghuy/talent-acquisition-management/pkg/container"
	"github.com/pghuy/talent-acquisition-management/pkg/http_response"
)

func (d *userDelivery) getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := auth.UIDFromContext(c)
		if uid == (auth.UID{}) {
			http_response.Response(c, http.StatusBadRequest, "invalid request", nil)
			return
		}

		http_response.Success(c, "success", container.Map{
			"id":       uid.ID,
			"username": uid.Username,
			"name":     uid.Username,
		})
	}
}
