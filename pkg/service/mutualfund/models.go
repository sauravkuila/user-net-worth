package mutualfund

type GetMutualFundsHoldingsResponse struct {
	Data  *GetMutualFundsHoldings `json:"data"`
	Error string                  `json:"error"`
}

type GetMutualFundsHoldings struct {
	InvestedValue float64     `json:"investedValue,omitempty"`
	CurrentValue  float64     `json:"currentValue,omitempty"`
	Holdings      interface{} `json:"holdings,omitempty"`
}

type UpdateHoldingsFromBrokerResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
