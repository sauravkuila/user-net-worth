package broker

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *brokerSt) UpdateHoldingsFromBroker(c *gin.Context) {
	var (
		request  UpdateHoldingsFromBrokerRequest
		response UpdateHoldingsFromBrokerResponse
	)

	if err := c.BindUri(&request); err != nil {
		log.Println("bad request", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if request.Broker == "angelone" {
		//fetch data from angel api
		holdings, err := external.GetHoldingsForAngel()
		if err != nil {
			log.Println("unable to fetch angel holdings", err.Error())
			response.Error = fmt.Sprintf("unable to fetch angel holdings. Error: %s", err.Error())
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		//insert holdings in db
		if err := obj.dbObj.InsertAngelOneHoldings(holdings); err != nil {
			log.Println("unable to update angel holdings", err.Error())
			response.Error = fmt.Sprintf("unable to update angel holdings. Error: %s", err.Error())
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		response.Data = "updated angelone holdings"
	} else if request.Broker == "zerodha" {
		holdings, err := external.GetHoldingsForZerodha()
		if err != nil {
			log.Println("unable to fetch idirect holdings", err.Error())
			response.Error = fmt.Sprintf("unable to fetch idirect holdings. Error: %s", err.Error())
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		//insert holdings in db
		if err := obj.dbObj.InsertZerodhaHoldings(holdings); err != nil {
			log.Println("unable to update idirect holdings", err.Error())
			response.Error = fmt.Sprintf("unable to update idirect holdings. Error: %s", err.Error())
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		response.Data = "updated zerodha holdings"
	} else if request.Broker == "idirect" {
		// holdings, err := external.GetHoldingsForIdirect()
		holdings, err := external.GetDematHoldingsForIDirect()
		if err != nil {
			log.Println("unable to fetch idirect holdings", err.Error())
			response.Error = fmt.Sprintf("unable to fetch idirect holdings. Error: %s", err.Error())
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		//insert holdings in db
		if err := obj.dbObj.InsertIDirectHoldings(holdings); err != nil {
			log.Println("unable to update idirect holdings", err.Error())
			response.Error = fmt.Sprintf("unable to update idirect holdings. Error: %s", err.Error())
			c.JSON(http.StatusFailedDependency, response)
			return
		}
		response.Data = "updated idirect holdings"
		response.Data = holdings
		// response.Error = "not implemented yet"
		// c.JSON(http.StatusPreconditionFailed, response)
		// return
	} else {
		response.Error = "invalid broker received"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (obj *brokerSt) GetSpecificBrokerHoldings(c *gin.Context) {
	var (
		request  GetSpecificBrokerHoldingsRequest
		response GetSpecificBrokerHoldingsResponse
	)
	log.Println("inside broker package")
	if err := c.BindUri(&request); err != nil {
		log.Println("bad request", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	holdings, investedVal, err := obj.getHoldingsByBroker(request.Broker)
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data = &GetSpecificBrokerHoldings{
		BrokerName:    request.Broker,
		Holdings:      holdings,
		InvestedValue: investedVal,
	}

	// if request.Broker == "angelone" {
	// 	holdings, investedVal, err := obj.dbObj.GetAngelOneHoldings()
	// 	if err != nil {
	// 		log.Println("failed to fetch holdings from db", err.Error())
	// 		response.Error = err.Error()
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	response.Data = &GetSpecificBrokerHoldings{
	// 		InvestedValue: investedVal,
	// 		Holdings:      holdings,
	// 	}
	// } else if request.Broker == "zerodha" {
	// 	holdings, investedVal, err := obj.dbObj.GetZerodhaHoldings()
	// 	if err != nil {
	// 		log.Println("failed to fetch holdings from db", err.Error())
	// 		response.Error = err.Error()
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	response.Data = &GetSpecificBrokerHoldings{
	// 		InvestedValue: investedVal,
	// 		Holdings:      holdings,
	// 	}
	// } else if request.Broker == "idirect" {
	// 	holdings, investedVal, err := obj.dbObj.GetIDirectHoldings()
	// 	if err != nil {
	// 		log.Println("failed to fetch holdings from db", err.Error())
	// 		response.Error = err.Error()
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	response.Data = &GetSpecificBrokerHoldings{
	// 		InvestedValue: investedVal,
	// 		Holdings:      holdings,
	// 	}
	// } else {
	// 	response.Error = "invalid broker received"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	c.JSON(http.StatusOK, response)
}