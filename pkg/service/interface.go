package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
)

type serviceStruct struct {
	dbObj db.DatabaseInterface
}

type ServiceInterface interface {
	GetSupportedBrokers(c *gin.Context)
	GetSpecificBrokerHoldings(c *gin.Context)
	GetAllBrokerHoldings() []HoldingsInfo
}

func InitService(dbItf db.DatabaseInterface) ServiceInterface {
	return &serviceStruct{
		dbObj: dbItf,
	}
}
