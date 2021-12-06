package talent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/talent-acquistion-management/auth"
	"github.com/pghuy/talent-acquistion-management/pkg/container"
	"github.com/pghuy/talent-acquistion-management/pkg/http_response"
)

func (d *talentDelivery) getUser() gin.HandlerFunc {
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
