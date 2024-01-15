package broker

import (
	"database/sql"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *brokerDbSt) GetAngelOneHoldings() ([]external.HoldingsInfo, float64, error) {
	var (
		investedValue float64 = 0
	)
	query := "select symbol, isin, quantity, price, updated_on from angel_one;"
	rows, err := obj.psql.Query(query)
	if err != nil {
		return nil, investedValue, err
	}
	holdings := make([]external.HoldingsInfo, 0)
	for rows.Next() {
		var (
			symbol   sql.NullString
			isin     sql.NullString
			quantity sql.NullInt64
			price    sql.NullFloat64
			updated  sql.NullTime
		)
		err := rows.Scan(&symbol, &isin, &quantity, &price, &updated)
		if err != nil {
			return nil, investedValue, err
		}
		holdings = append(holdings, external.HoldingsInfo{
			Symbol:    symbol.String,
			Quantity:  quantity.Int64,
			Isin:      isin.String,
			AvgPrice:  price.Float64,
			UpdatedOn: updated.Time,
		})
		investedValue += float64(quantity.Int64) * price.Float64
	}
	return holdings, investedValue, nil
}

func (obj *brokerDbSt) GetIDirectHoldings() ([]external.HoldingsInfo, float64, error) {
	var (
		investedValue float64 = 0
	)
	query := "select symbol, isin, quantity, price, updated_on from idirect;"
	rows, err := obj.psql.Query(query)
	if err != nil {
		return nil, investedValue, err
	}
	holdings := make([]external.HoldingsInfo, 0)
	for rows.Next() {
		var (
			symbol   sql.NullString
			isin     sql.NullString
			quantity sql.NullInt64
			price    sql.NullFloat64
			updated  sql.NullTime
		)
		err := rows.Scan(&symbol, &isin, &quantity, &price, &updated)
		if err != nil {
			return nil, investedValue, err
		}
		holdings = append(holdings, external.HoldingsInfo{
			Symbol:    symbol.String,
			Quantity:  quantity.Int64,
			Isin:      isin.String,
			AvgPrice:  price.Float64,
			UpdatedOn: updated.Time,
		})
		investedValue += float64(quantity.Int64) * price.Float64
	}
	return holdings, investedValue, nil
}

func (obj *brokerDbSt) GetZerodhaHoldings() ([]external.HoldingsInfo, float64, error) {
	var (
		investedValue float64 = 0
	)
	query := "select symbol, isin, quantity, price, updated_on from zerodha;"
	rows, err := obj.psql.Query(query)
	if err != nil {
		return nil, investedValue, err
	}
	holdings := make([]external.HoldingsInfo, 0)
	for rows.Next() {
		var (
			symbol   sql.NullString
			isin     sql.NullString
			quantity sql.NullInt64
			price    sql.NullFloat64
			updated  sql.NullTime
		)
		err := rows.Scan(&symbol, &isin, &quantity, &price, &updated)
		if err != nil {
			return nil, investedValue, err
		}
		holdings = append(holdings, external.HoldingsInfo{
			Symbol:    symbol.String,
			Quantity:  quantity.Int64,
			Isin:      isin.String,
			AvgPrice:  price.Float64,
			UpdatedOn: updated.Time,
		})
		investedValue += float64(quantity.Int64) * price.Float64
	}
	return holdings, investedValue, nil
}
