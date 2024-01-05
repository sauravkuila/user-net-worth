package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/service"
)

func getRouter(serviceObj service.ServiceInterface) *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I am Healthy")
	})
	router.GET("/brokers", serviceObj.GetSupportedBrokers)
	router.GET("/holding/:broker", serviceObj.GetSpecificBrokerHoldings)

	return router
}
