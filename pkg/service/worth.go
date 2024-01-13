package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, response)
}
