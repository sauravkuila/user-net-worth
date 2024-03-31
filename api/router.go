package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/service"
)

func getRouter(serviceObj service.ServiceInterface) *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I am Healthy")
	})
	router.GET("/checkDb", func(c *gin.Context) {
		_, err := db.InitDb()
		c.JSON(http.StatusOK, gin.H{"error": err})
	})
	//sources for the data
	sourceGroup := router.Group("/sources")
	{
		sourceGroup.GET("", serviceObj.GetCredsObject().GetSupportedSources)
		sourceGroup.POST("/broker/sync", serviceObj.GetCredsObject().UpdateBrokerCred)
		sourceGroup.POST("/mutualfund/sync", serviceObj.GetCredsObject().UpdateMutualFundCred)
	}

	//holding against relevant sources
	holdingGroup := router.Group("/holding")
	{
		holdingGroup.GET("/broker/sync/:broker", serviceObj.GetBrokerObject().UpdateHoldingsFromBroker)
		holdingGroup.GET("/broker/:broker", serviceObj.GetBrokerObject().GetSpecificBrokerHoldings)
		holdingGroup.GET("/mf/sync", serviceObj.GetMutualFundObject().UpdateMfHoldingsFromMfCentral)
		holdingGroup.GET("/mf", serviceObj.GetMutualFundObject().GetMutualFundsHoldings)
	}

	//net worth
	router.GET("/worth", serviceObj.GetTotalWorth)

	//callback for source dev apps
	callbackGroup := router.Group("/callback")
	{
		// callbackGroup.GET("/idirect", serviceObj.GetCallbackObject().CallbackForIdirectApiSessionKey)
		callbackGroup.POST("/idirect", serviceObj.GetCallbackObject().CallbackForIdirectApiSessionKey) //using this
	}

	return router
}
