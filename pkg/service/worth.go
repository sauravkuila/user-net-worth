package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/service/broker"
	"github.com/sauravkuila/portfolio-worth/pkg/service/mutualfund"
)

func (obj *serviceStruct) GetTotalWorth(c *gin.Context) {
	var (
		response GetTotalWorthResponse
	)

	response.Data = &GetTotalWorth{
		Stocks: make([]broker.GetSpecificBrokerHoldings, 0),
	}
	//broker holdings
	brokerHoldings, err := obj.brokerObj.GetAllBrokerHoldings()
	if err != nil {
		log.Println("error fetching holdings", err.Error())
		c.JSON(http.StatusFailedDependency, response)
		return
	}
	for _, v := range brokerHoldings {
		response.Data.Stocks = append(response.Data.Stocks, v)
		response.Data.TotalInvestedValue += v.InvestedValue
		response.Data.TotalCurrentValue += v.CurrentValue
	}

	//mutual funds
	{
		holdings, investedVal, currentVal, err := obj.mfObj.GetMutualFundHoldingData()
		if err != nil {
			log.Println("failed to fetch holdings from db", err.Error())
			response.Error = err.Error()
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		response.Data.MutualFunds = &mutualfund.GetMutualFundsHoldings{
			InvestedValue: investedVal,
			CurrentValue:  currentVal,
			Holdings:      holdings,
		}
		response.Data.TotalInvestedValue += response.Data.MutualFunds.InvestedValue
		response.Data.TotalCurrentValue += response.Data.MutualFunds.CurrentValue
	}

	c.JSON(http.StatusOK, response)
}
