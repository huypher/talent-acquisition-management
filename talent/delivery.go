package talent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huypher/talent-acquisition-management/domain"
)

type talentDelivery struct {
	talentUsecase domain.TalentUsecase
}

func NewTalentDelivery(userUsecase domain.TalentUsecase) *talentDelivery {
	return &talentDelivery{
		talentUsecase: userUsecase,
	}
}

func (d *talentDelivery) Handler(c *gin.RouterGroup) {
	c.Handle(http.MethodGet, "/talents", d.getTalentList())
}
