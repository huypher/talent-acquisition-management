package talent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/dobi-oms/domain"
)

type talentDelivery struct {
	userUsecase domain.TalentUsecase
}

func NewTalentDelivery(userUsecase domain.TalentUsecase) *talentDelivery {
	return &talentDelivery{
		userUsecase: userUsecase,
	}
}

func (d *talentDelivery) Handler(c *gin.RouterGroup) {
	c.Handle(http.MethodGet, "/talent", d.getUser())
}
