package external

import "time"

type HoldingsInfo struct {
	Symbol    string
	Quantity  int64
	Isin      string
	AvgPrice  float64
	UpdatedOn time.Time
}
