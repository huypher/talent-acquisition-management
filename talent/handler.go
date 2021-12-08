package talent

import (
	"github.com/gin-gonic/gin"
	"github.com/huypher/utils/container"
	"github.com/huypher/utils/http_response"
	"github.com/huypher/utils/transformer"
)

func (d *talentDelivery) getTalentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getTalentList getTalentList

		if err := c.ShouldBindJSON(&getTalentList); err != nil {
			http_response.BadRequest(c, err.Error(), nil)
			return
		}

		params := transformer.FlattenStructToContainerMap(&getTalentList)
		filter := params.Exclude([]string{"page_id", "per_page"})

		pageID := getTalentList.PageID
		perPage := getTalentList.PerPage

		talents, err := d.talentUsecase.GetList(c, filter, pageID, perPage)
		if err != nil {
			http_response.Error(c, err)
		}

		http_response.Success(c, "success", container.Map{
			"talents": talents,
		})
	}
}

type getTalentList struct {
	FullName string `json:"full_name"`
	PageID   int    `json:"page_id"`
	PerPage  int    `json:"per_page"`
}
