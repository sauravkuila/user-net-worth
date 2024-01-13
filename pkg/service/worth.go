package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

func (obj *serviceStruct) GetTotalWorth(c *gin.Context) {
	var (
		response GetTotalWorthResponse
	)

	response.Data = &GetTotalWorth{
		Stocks: make([]GetSpecificBrokerHoldings, 0),
	}

	holdings, investedVal, err := obj.dbObj.GetAngelOneHoldings()
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data.TotalInvested += investedVal
	response.Data.Stocks = append(response.Data.Stocks, GetSpecificBrokerHoldings{
		BrokerName:    "AngelOne",
		InvestedValue: investedVal,
		Holdings:      holdings,
	})

	holdings, investedVal, err = obj.dbObj.GetZerodhaHoldings()
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data.TotalInvested += investedVal
	response.Data.Stocks = append(response.Data.Stocks, GetSpecificBrokerHoldings{
		BrokerName:    "Zerodha",
		InvestedValue: investedVal,
		Holdings:      holdings,
	})

	holdings, investedVal, err = obj.dbObj.GetIDirectHoldings()
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data.TotalInvested += investedVal
	response.Data.Stocks = append(response.Data.Stocks, GetSpecificBrokerHoldings{
		BrokerName:    "IDirect",
		InvestedValue: investedVal,
		Holdings:      holdings,
	})

	//mutual funds
	{
		mfholdings, mfInvestedVal, err := obj.dbObj.GetMfCentralHoldings()
		if err != nil {
			log.Println("failed to fetch holdings from db", err.Error())
			response.Error = err.Error()
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response.Data.TotalInvested += mfInvestedVal
		response.Data.MutualFunds = &GetMutualFundsHoldings{
			InvestedValue: mfInvestedVal,
			// Holdings:      mfholdings,
		}
		nonZeroFolios := make([]external.MfHoldingsInfo, 0)

		// ignore zeo value folios
		for _, holding := range mfholdings {
			if holding.Quantity <= 0 {
				continue
			}
			currVal := holding.CurrentVal
			nav, err := utils.GetNavValueFromIsin(holding.Isin)
			if err == nil {
				currVal = nav * holding.Quantity
			}
			response.Data.MutualFunds.CurrentValue += currVal
			nonZeroFolios = append(nonZeroFolios, holding)
		}
		response.Data.MutualFunds.Holdings = nonZeroFolios
	}

	c.JSON(http.StatusOK, response)
}
