package broker

import (
	"database/sql"

	"github.com/sauravkuila/portfolio-worth/external"
)

type brokerDbSt struct {
	psql *sql.DB
}

type BrokerDatabaseInterface interface {
	GetAngelOneHoldings() ([]external.HoldingsInfo, float64, error)
	GetIDirectHoldings() ([]external.HoldingsInfo, float64, error)
	GetZerodhaHoldings() ([]external.HoldingsInfo, float64, error)
	InsertAngelOneHoldings([]external.HoldingsInfo) error
	InsertIDirectHoldings([]external.HoldingsInfo) error
	InsertZerodhaHoldings([]external.HoldingsInfo) error
}

func NewBrokerDBInterface(db *sql.DB) BrokerDatabaseInterface {
	return &brokerDbSt{
		psql: db,
	}
}
