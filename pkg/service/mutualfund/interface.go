package mutualfund

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
)

type mutualfundSt struct {
	dbObj db.DatabaseInterface
}

type MutualFundInterface interface {
	GetMutualFundsHoldings(c *gin.Context)
	UpdateMfHoldingsFromMfCentral(c *gin.Context)
}

func NewMutualfundInterfaceObj(db db.DatabaseInterface) MutualFundInterface {
	return &mutualfundSt{
		dbObj: db,
	}
}
