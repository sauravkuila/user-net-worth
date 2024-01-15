package broker

type UpdateHoldingsFromBrokerRequest struct {
	Broker string `uri:"broker" binding:"required"`
}

type UpdateHoldingsFromBrokerResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type GetSpecificBrokerHoldingsRequest struct {
	Broker string `uri:"broker" binding:"required"`
}

type GetSpecificBrokerHoldingsResponse struct {
	Data  *GetSpecificBrokerHoldings `json:"data,omitempty"`
	Error string                     `json:"error,omitempty"`
}

type GetSpecificBrokerHoldings struct {
	BrokerName    string      `json:"broker,omitempty"`
	InvestedValue float64     `json:"investedValue,omitempty"`
	Holdings      interface{} `json:"holdings,omitempty"`
}
