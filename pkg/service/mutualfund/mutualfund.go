package mutualfund

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

func (obj *mutualfundSt) GetMutualFundsHoldings(c *gin.Context) {
	var (
		response GetMutualFundsHoldingsResponse
	)

	//fetch data from db
	holdings, investedVal, err := obj.dbObj.GetMfCentralHoldings()
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data = &GetMutualFundsHoldings{
		InvestedValue: investedVal,
		Holdings:      holdings,
	}

	for _, holding := range holdings {
		currVal := holding.CurrentVal
		nav, err := utils.GetNavValueFromIsin(holding.Isin)
		if err == nil {
			currVal = nav * holding.Quantity
		}
		response.Data.CurrentValue += currVal
	}

	c.JSON(http.StatusOK, response)
}

func (obj *mutualfundSt) UpdateMfHoldingsFromMfCentral(c *gin.Context) {
	var (
		response UpdateHoldingsFromBrokerResponse
	)

	//fetch data from mf central
	holdings, err := external.GetMutualFundHoldingsFromMfCentral()
	if err != nil {
		log.Println("unable to fetch angel holdings", err.Error())
		response.Error = fmt.Sprintf("unable to fetch angel holdings. Error: %s", err.Error())
		c.JSON(http.StatusFailedDependency, response)
		return
	}
	//insert holdings in db
	if err := obj.dbObj.InsertMfCentralHoldings(holdings); err != nil {
		log.Println("unable to update angel holdings", err.Error())
		response.Error = fmt.Sprintf("unable to update angel holdings. Error: %s", err.Error())
		c.JSON(http.StatusFailedDependency, response)
		return
	}
	// response.Data = "updated mutual fund holdings"
	response.Data = holdings

	c.JSON(http.StatusOK, response)
}
