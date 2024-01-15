package service

import (
	"log"

	"github.com/sauravkuila/portfolio-worth/pkg/service/broker"
	"github.com/sauravkuila/portfolio-worth/pkg/service/callback"
	"github.com/sauravkuila/portfolio-worth/pkg/service/creds"
	"github.com/sauravkuila/portfolio-worth/pkg/service/mutualfund"
)

func (obj *serviceStruct) GetBrokerObject() broker.BrokerInterface {
	log.Println("fetched broker Interface")
	return obj.brokerObj
}

func (obj *serviceStruct) GetMutualFundObject() mutualfund.MutualFundInterface {
	log.Println("fetched mutualfund Interface")
	return obj.mfObj
}

func (obj *serviceStruct) GetCredsObject() creds.CredsInterface {
	log.Println("fetched cred Interface")
	return obj.credObj
}

func (obj *serviceStruct) GetCallbackObject() callback.CallbackInterface {
	log.Println("fetched callback Interface")
	return obj.callbackObj
}
