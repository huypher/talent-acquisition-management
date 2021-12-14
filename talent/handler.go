package talent

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huypher/kit/container"
	"github.com/huypher/kit/http_response"
	"github.com/huypher/kit/utils"
	"github.com/huypher/talent-acquisition-management/domain"
	"github.com/sirupsen/logrus"
)

func (d *talentDelivery) getTalentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r getTalentListRequest

		if err := c.ShouldBindQuery(&r); err != nil {
			logrus.Errorf("bind request body err=%v\n", err)
			http_response.BadRequest(c, err.Error(), nil)
			return
		}

		params := utils.FlattenStructToContainerMap(&r)
		filter := params.Exclude([]string{"page_id", "per_page"})

		pageID := r.PageID
		perPage := r.PerPage

		talents, err := d.talentUsecase.GetList(c, filter, pageID, perPage)
		if err != nil {
			http_response.Error(c, err)
		}

		http_response.Success(c, "success", container.Map{
			"talents": talents,
		})
	}
}

func (d *talentDelivery) addTalent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r addTalentRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			http_response.BadRequest(c, err.Error(), nil)
			return
		}

		model := domain.Talent{
			FullName:           r.FullName,
			Gender:             r.Gender,
			Birthdate:          r.Birthdate,
			Phone:              r.Phone,
			Email:              r.Email,
			AppliedPosition:    r.AppliedPosition,
			Level:              r.Level,
			Department:         r.Department,
			Project:            r.Project,
			CV:                 r.CV,
			Criteria:           r.Criteria,
			ScheduledInterview: r.ScheduledInterview,
			InterviewResult:    r.InterviewResult,
		}

		err := d.talentUsecase.AddTalent(c, model)
		if err != nil {
			http_response.Error(c, err)
			return
		}

		http_response.Success(c, "success", nil)
	}
}

func (d *talentDelivery) updateTalent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r updateTalentRequest

		talentIDInput := c.Param("talent_id")
		if talentIDInput == "" {
			http_response.BadRequest(c, "invalid request", nil)
			return
		}

		talentID, err := strconv.Atoi(talentIDInput)
		if err != nil {
			http_response.BadRequest(c, err.Error(), nil)
			return
		}
		if talentID <= 0 {
			http_response.BadRequest(c, "invalid request", nil)
			return
		}

		if err := c.ShouldBindJSON(&r); err != nil {
			http_response.BadRequest(c, err.Error(), nil)
			return
		}

		params := utils.FlattenStructToContainerMap(&r)

		err = d.talentUsecase.UpdateTalent(c, talentID, params)
		if err != nil {
			var errTalentNotfound *ErrTalentNotFound
			switch {
			case errors.As(err, &errTalentNotfound):
				http_response.BadRequest(c, err.Error(), nil)
				return
			default:
				http_response.Error(c, err)
				return
			}
		}

		http_response.Success(c, "success", nil)
	}
}
