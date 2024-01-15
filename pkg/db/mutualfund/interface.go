package mutualfund

import (
	"database/sql"

	"github.com/sauravkuila/portfolio-worth/external"
)

type mutualfundDbSt struct {
	psql *sql.DB
}

type MutualFundDatabaseInterface interface {
	GetMfCentralHoldings() ([]external.MfHoldingsInfo, float64, error)
	InsertMfCentralHoldings([]external.MfHoldingsInfo) error
}

func NewMutualfundInterfaceObj(db *sql.DB) MutualFundDatabaseInterface {
	return &mutualfundDbSt{
		psql: db,
	}
}
