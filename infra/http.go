package infra

import (
	"net/http"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/huypher/talent-acquisition-management/auth"
	"github.com/huypher/talent-acquisition-management/domain"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler(
	authDelivery domain.AuthDelivery,
	talentDelivery domain.TalentDelivery,
) http.Handler {
	router := gin.Default()

	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(healthcheck.Default())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	v1 := router.Group("/v1")
	{
		authDelivery.Handler(v1)
		v1.Use(auth.Middleware())
		talentDelivery.Handler(v1)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
