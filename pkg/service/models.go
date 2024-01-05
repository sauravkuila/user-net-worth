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
