package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *serviceStruct) GetSupportedBrokers(c *gin.Context) {
	//get supported brokers from db
	// brokers := []string{"AngelOne", "Zerodha", "ICICIDirect"}
	var brokers []string

	brokers, err := obj.dbObj.GetSupportedBrokers()
	if err != nil {
		log.Println("unable to fetch supported brokers", err.Error())
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"data": brokers,
	})
	return
}

func (obj *serviceStruct) GetSpecificBrokerHoldings(c *gin.Context) {
	var (
		request  GetSpecificBrokerHoldingsRequest
		response GetSpecificBrokerHoldingsResponse
		err      error
	)
	if err := c.BindUri(&request); err != nil {
		log.Println("bad request", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if request.Broker == "angelone" {
		holdings, lastUpdated, investedVal, _ := obj.dbObj.GetAngelOneHoldings()
		if (len(holdings) != 0 && lastUpdated.YearDay() < time.Now().YearDay()) || len(holdings) == 0 {
			//fetch data from angel api
			holdings, err = external.GetHoldingsForAngel()
			if err != nil {
				log.Println("unable to fetch angel holdings", err.Error())
				response.Error = err.Error()
				c.JSON(http.StatusFailedDependency, response)
				return
			}
			//insert holdings in db
			if err := obj.dbObj.InsertAngelOneHoldings(holdings); err != nil {
				log.Println("unable to update angel holdings", err.Error())
			}
		}
		response.Data = &GetSpecificBrokerHoldings{
			InvestedValue: investedVal,
			Holdings:      holdings,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	response.Error = "invalid broker received"
	c.JSON(http.StatusBadRequest, response)
}

func (obj *serviceStruct) GetAllBrokerHoldings() []HoldingsInfo {
	return nil
}
