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
	router.POST("/brokers/sync", serviceObj.UpdateBrokerCred)
	router.GET("/holding/sync/:broker", serviceObj.UpdateHoldingsFromBroker)
	router.GET("/holding/:broker", serviceObj.GetSpecificBrokerHoldings)
	router.GET("/holding/mf", serviceObj.GetMutualFundsHoldings)
	router.GET("/holding/mf/sync", serviceObj.UpdateMfHoldingsFromMfCentral)
	router.GET("/worth", serviceObj.GetTotalWorth)

	callbackGroup := router.Group("/callback")
	{
		callbackGroup.GET("/idirect", serviceObj.GetIdirectApiSessionKey)
		callbackGroup.POST("/idirect", serviceObj.GetIdirectApiSessionKey) //using this
	}

	return router
}
