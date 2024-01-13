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
	BrokerName    string      `json:"broker,omitempty"`
	InvestedValue float64     `json:"investedValue,omitempty"`
	Holdings      interface{} `json:"holdings,omitempty"`
}

type GetIdirectSessionKeyRequest struct {
	ApiSession int64 `form:"apisession" binding:"required"`
}

type UpdateHoldingsFromBrokerRequest struct {
	Broker string `uri:"broker" binding:"required"`
}

type UpdateHoldingsFromBrokerResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type GetTotalWorthResponse struct {
	Data  *GetTotalWorth `json:"data,omitempty"`
	Error string         `json:"error,omitempty"`
}

type GetTotalWorth struct {
	TotalInvested float64                     `json:"totalInvested,omitempty"`
	Stocks        []GetSpecificBrokerHoldings `json:"stocks,omitempty"`
}

type UpdateBrokerCredRequest struct {
	Broker     string `uri:"broker" binding:"required"`
	TOTPSecret string `json:"totp_secret"`
	UserKey    string `json:"user_key"`
	PassKey    string `json:"pass_key"`
	AppCode    string `json:"app_code"`
	SecretKey  string `json:"secret_key"`
}

type UpdateBrokerCredResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
