package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/service/broker"
	"github.com/sauravkuila/portfolio-worth/pkg/service/callback"
	"github.com/sauravkuila/portfolio-worth/pkg/service/creds"
	"github.com/sauravkuila/portfolio-worth/pkg/service/mutualfund"
)

type serviceStruct struct {
	dbObj       db.DatabaseInterface
	brokerObj   broker.BrokerInterface
	mfObj       mutualfund.MutualFundInterface
	credObj     creds.CredsInterface
	callbackObj callback.CallbackInterface
}

type ServiceInterface interface {
	GetCredsObject() creds.CredsInterface
	GetBrokerObject() broker.BrokerInterface
	GetMutualFundObject() mutualfund.MutualFundInterface
	GetCallbackObject() callback.CallbackInterface
	GetTotalWorth(c *gin.Context)
}

func InitService(dbItf db.DatabaseInterface) ServiceInterface {
	serviceObj := serviceStruct{
		dbObj:       dbItf,
		brokerObj:   broker.NewBrokerInterfaceObj(dbItf),
		mfObj:       mutualfund.NewMutualfundInterfaceObj(dbItf),
		credObj:     creds.NewCredsInterfaceObject(dbItf),
		callbackObj: callback.NewCallbackInterface(dbItf),
	}
	// go serviceObj.initBrokerLogins()
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
