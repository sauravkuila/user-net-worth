package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
)

type brokerSt struct {
	dbObj db.DatabaseInterface
}

type BrokerInterface interface {
	UpdateHoldingsFromBroker(c *gin.Context)
	GetSpecificBrokerHoldings(c *gin.Context)
	//returns all holdings across brokers
	//output: holdings
	GetAllBrokerHoldings() (map[string]GetSpecificBrokerHoldings, error)
}

func NewBrokerInterfaceObj(db db.DatabaseInterface) BrokerInterface {
	return &brokerSt{
		dbObj: db,
	}
}
