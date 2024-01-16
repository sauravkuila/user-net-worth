package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/quote"
)

type brokerSt struct {
	dbObj    db.DatabaseInterface
	quoteObj quote.QuoteInterface
}

type BrokerInterface interface {
	UpdateHoldingsFromBroker(c *gin.Context)
	GetSpecificBrokerHoldings(c *gin.Context)
	//returns all holdings across brokers
	//	output:
	//		map[broker]HoldingInfo objec
	//		error
	GetAllBrokerHoldings() (map[string]GetSpecificBrokerHoldings, error)
	//returns all token across all holdings including all brokers
	GetHoldingTokenAcrossAllBrokers() ([]string, error)
}

func NewBrokerInterfaceObj(db db.DatabaseInterface, quoteItf quote.QuoteInterface) BrokerInterface {
	brokerObj := &brokerSt{
		dbObj:    db,
		quoteObj: quoteItf,
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
	if creds, err := obj.dbObj.GetBrokerCred("idirect"); err == nil {
		external.LoginAndSyncIDirect(creds["app_api_key"].(string), creds["secret_key"].(string))
	}
}

// alphavantage api_key
// I1ZZ2MZHEUSB9Q5O
