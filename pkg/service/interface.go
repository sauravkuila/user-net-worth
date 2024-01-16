package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/quote"
	"github.com/sauravkuila/portfolio-worth/pkg/service/broker"
	"github.com/sauravkuila/portfolio-worth/pkg/service/callback"
	"github.com/sauravkuila/portfolio-worth/pkg/service/creds"
	"github.com/sauravkuila/portfolio-worth/pkg/service/mutualfund"
)

type serviceStruct struct {
	dbObj       db.DatabaseInterface
	quoteObj    quote.QuoteInterface
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

func InitService(dbItf db.DatabaseInterface, quoteItf quote.QuoteInterface) ServiceInterface {
	serviceObj := serviceStruct{
		dbObj:       dbItf,
		quoteObj:    quoteItf,
		brokerObj:   broker.NewBrokerInterfaceObj(dbItf, quoteItf),
		mfObj:       mutualfund.NewMutualfundInterfaceObj(dbItf, quoteItf),
		credObj:     creds.NewCredsInterfaceObject(dbItf),
		callbackObj: callback.NewCallbackInterface(dbItf),
	}
	return &serviceObj
}
