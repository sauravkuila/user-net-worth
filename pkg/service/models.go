package service

import (
	"github.com/sauravkuila/portfolio-worth/pkg/service/broker"
	"github.com/sauravkuila/portfolio-worth/pkg/service/mutualfund"
)

type GetTotalWorthResponse struct {
	Data  *GetTotalWorth `json:"data,omitempty"`
	Error string         `json:"error,omitempty"`
}

type GetTotalWorth struct {
	TotalInvestedValue float64                            `json:"totalInvestedValue,omitempty"`
	TotalCurrentValue  float64                            `json:"totalCurrentValue,omitempty"`
	Stocks             []broker.GetSpecificBrokerHoldings `json:"stocks,omitempty"`
	MutualFunds        *mutualfund.GetMutualFundsHoldings `json:"mutualfunds,omitempty"`
}
