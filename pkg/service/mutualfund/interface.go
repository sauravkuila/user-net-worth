package mutualfund

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/quote"
)

type mutualfundSt struct {
	dbObj    db.DatabaseInterface
	quoteObj quote.QuoteInterface
}

type MutualFundInterface interface {
	GetMutualFundsHoldings(c *gin.Context)
	UpdateMfHoldingsFromMfCentral(c *gin.Context)
	GetMutualFundHoldingData() ([]external.MfHoldingsInfo, float64, float64, error)
	GetIsinOfFundHoldings() ([]string, error)
}

func NewMutualfundInterfaceObj(db db.DatabaseInterface, quoteItf quote.QuoteInterface) MutualFundInterface {
	return &mutualfundSt{
		dbObj:    db,
		quoteObj: quoteItf,
	}
}
