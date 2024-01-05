package service

type HoldingsInfo struct {
	Symbol   string
	Quantity int64
	Isin     string
	AvgPrice float64
}

type GetSpecificBrokerHoldingsRequest struct {
	Broker string `uri:"broker" binding:"required"`
}

type GetSpecificBrokerHoldingsResponse struct {
	Data  *GetSpecificBrokerHoldings `json:"data,omitempty"`
	Error string                     `json:"error,omitempty"`
}

type GetSpecificBrokerHoldings struct {
	InvestedValue float64     `json:"investedValue,omitempty"`
	Holdings      interface{} `json:"holdings,omitempty"`
}
