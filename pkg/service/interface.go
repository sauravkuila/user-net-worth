package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
)

type serviceStruct struct {
	dbObj db.DatabaseInterface
}

type ServiceInterface interface {
	GetSupportedBrokers(c *gin.Context)
	UpdateBrokerCred(c *gin.Context)

	UpdateHoldingsFromBroker(c *gin.Context)
	GetSpecificBrokerHoldings(c *gin.Context)
	GetAllBrokerHoldings() []HoldingsInfo

	GetMutualFundsHoldings(c *gin.Context)
	UpdateMfHoldingsFromMfCentral(c *gin.Context)

	GetIdirectApiSessionKey(c *gin.Context)
	GetTotalWorth(c *gin.Context)
}

func InitService(dbItf db.DatabaseInterface) ServiceInterface {
	serviceObj := serviceStruct{
		dbObj: dbItf,
	}
	go serviceObj.initBrokerLogins()
	return &serviceObj
}

func (obj *serviceStruct) initBrokerLogins() {
	//login all brokers to generate sessions
	if creds, err := obj.dbObj.GetBrokerCred("angelone"); err == nil {
		external.LoginAndSyncAngelOne(creds["user_key"].(string), creds["pass_key"].(string), creds["totp_secret"].(string), creds["app_api_key"].(string))
	} else {
		log.Println("failed")
	}
	if creds, err := obj.dbObj.GetBrokerCred("zerodha"); err == nil {
		external.LoginAndSyncZerodha(creds["user_key"].(string), creds["pass_key"].(string), creds["totp_secret"].(string))
	}
}

// alphavantage api_key
// I1ZZ2MZHEUSB9Q5O
