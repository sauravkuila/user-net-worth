package mutualfund

import (
	"database/sql"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *mutualfundDbSt) GetMfCentralHoldings() ([]external.MfHoldingsInfo, float64, error) {
	var (
		investedValue float64 = 0
	)
	query := "select folio, scheme_name, isin, quantity, price, cost_price, curr_price, updated_on from mf_central;"
	rows, err := obj.psql.Query(query)
	if err != nil {
		return nil, investedValue, err
	}
	holdings := make([]external.MfHoldingsInfo, 0)
	for rows.Next() {
		var (
			folio       sql.NullString
			schemeName  sql.NullString
			isin        sql.NullString
			quantity    sql.NullFloat64
			price       sql.NullFloat64
			investedVal sql.NullFloat64
			currentVal  sql.NullFloat64
			updated     sql.NullTime
		)
		err := rows.Scan(&folio, &schemeName, &isin, &quantity, &price, &investedVal, &currentVal, &updated)
		if err != nil {
			return nil, investedValue, err
		}
		holdings = append(holdings, external.MfHoldingsInfo{
			Folio:       folio.String,
			Name:        schemeName.String,
			Quantity:    quantity.Float64,
			Isin:        isin.String,
			AvgPrice:    price.Float64,
			InvestedVal: investedVal.Float64,
			CurrentVal:  currentVal.Float64,
			UpdatedOn:   updated.Time,
		})
		investedValue += investedVal.Float64
	}
	return holdings, investedValue, nil
}
