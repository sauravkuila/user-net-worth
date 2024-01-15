package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
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
	brokerObj := &brokerSt{
		dbObj: db,
	}
	go brokerObj.initBrokerLogins()
	return brokerObj
}

func (obj *brokerSt) initBrokerLogins() {
	//login all brokers to generate sessions
	if creds, err := obj.dbObj.GetBrokerCred("angelone"); err == nil {
		external.LoginAndSyncAngelOne(creds["user_key"].(string), creds["pass_key"].(string), creds["totp_secret"].(string), creds["app_api_key"].(string))
	}
	if creds, err := obj.dbObj.GetBrokerCred("zerodha"); err == nil {
		external.LoginAndSyncZerodha(creds["user_key"].(string), creds["pass_key"].(string), creds["totp_secret"].(string))
	}
}

// alphavantage api_key
// I1ZZ2MZHEUSB9Q5O
