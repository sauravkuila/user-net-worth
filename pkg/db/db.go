package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *databaseStruct) GetSupportedBrokers() ([]string, error) {
	query := "select broker_name from supported_broker;"
	rows, err := obj.psql.Query(query)
	if err != nil {
		return nil, err
	}
	brokers := make([]string, 0)
	for rows.Next() {
		var broker sql.NullString
		err := rows.Scan(&broker)
		if err != nil {
			return nil, err
		}
		brokers = append(brokers, broker.String)
	}
	return brokers, nil
}

func (obj *databaseStruct) GetAngelOneHoldings() ([]external.HoldingsInfo, float64, error) {
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

func (obj *databaseStruct) GetIDirectHoldings() ([]external.HoldingsInfo, float64, error) {
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

func (obj *databaseStruct) GetZerodhaHoldings() ([]external.HoldingsInfo, float64, error) {
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

func (obj *databaseStruct) InsertAngelOneHoldings(holdings []external.HoldingsInfo) error {
	//flush db
	tx, err := obj.psql.Begin()
	if err != nil {
		log.Println("txn error. ", err.Error())
		return err
	}
	delQuery := "delete from angel_one;"
	_, err = tx.Exec(delQuery)
	if err != nil {
		log.Println("failed to flush. ", err.Error())
		tx.Rollback()
		return err
	}
	updateBrokerStatus := "update supported_broker set holdings_sync='t' where broker_name='angelone';"
	res, err := tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error in updation", err.Error())
		tx.Rollback()
		return err
	}
	log.Println(affected)

	query := "insert into angel_one(symbol, isin, quantity, price) values "
	var args []interface{}
	sQueries := make([]string, 0)
	for i, holding := range holdings {
		// q := "(?,?,?,?)"
		q := fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4)
		sQueries = append(sQueries, q)
		args = append(args, holding.Symbol, holding.Isin, holding.Quantity, holding.AvgPrice)
	}

	query += strings.Join(sQueries, ",") + ";"

	// res, err := obj.psql.Exec(query, args...)
	res, err = tx.Exec(query, args...)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err = res.RowsAffected()
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	if int(affected) != len(holdings) {
		tx.Rollback()
		return errors.New("not all holdings were inserted")
	}

	updateBrokerStatus = "update supported_broker set holdings_sync='t' where broker_name='angelone';"
	res, err = tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err = res.RowsAffected()
	if err != nil {
		log.Println("Error in updation", err.Error())
		tx.Rollback()
		return err
	}
	log.Println(affected)

	return tx.Commit()
}

func (obj *databaseStruct) InsertIDirectHoldings(holdings []external.HoldingsInfo) error {
	//flush db
	tx, err := obj.psql.Begin()
	if err != nil {
		log.Println("txn error. ", err.Error())
		return err
	}
	delQuery := "delete from idirect;"
	_, err = tx.Exec(delQuery)
	if err != nil {
		log.Println("failed to flush. ", err.Error())
		tx.Rollback()
		return err
	}

	updateBrokerStatus := "update supported_broker set holdings_sync='f' where broker_name='idirect';"
	res, err := tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error in updation", err.Error())
		tx.Rollback()
		return err
	}
	log.Println(affected)

	query := "insert into idirect(symbol, isin, quantity, price) values "
	var args []interface{}
	sQueries := make([]string, 0)
	for i, holding := range holdings {
		// q := "(?,?,?,?)"
		q := fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4)
		sQueries = append(sQueries, q)
		args = append(args, holding.Symbol, holding.Isin, holding.Quantity, holding.AvgPrice)
	}

	query += strings.Join(sQueries, ",") + ";"

	// res, err := obj.psql.Exec(query, args...)
	res, err = tx.Exec(query, args...)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err = res.RowsAffected()
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	if int(affected) != len(holdings) {
		tx.Rollback()
		return errors.New("not all holdings were inserted")
	}

	updateBrokerStatus = "update supported_broker set holdings_sync='t' where broker_name='idirect';"
	res, err = tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err = res.RowsAffected()
	if err != nil {
		log.Println("Error in updation", err.Error())
		tx.Rollback()
		return err
	}
	log.Println(affected)

	return tx.Commit()
}

func (obj *databaseStruct) InsertZerodhaHoldings(holdings []external.HoldingsInfo) error {
	//flush db
	tx, err := obj.psql.Begin()
	if err != nil {
		log.Println("txn error. ", err.Error())
		return err
	}
	delQuery := "delete from zerodha;"
	_, err = tx.Exec(delQuery)
	if err != nil {
		log.Println("failed to flush. ", err.Error())
		tx.Rollback()
		return err
	}

	updateBrokerStatus := "update supported_broker set holdings_sync='f' where broker_name='zerodha';"
	res, err := tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error in updation", err.Error())
		tx.Rollback()
		return err
	}
	log.Println(affected)

	query := "insert into zerodha(symbol, isin, quantity, price) values "
	var args []interface{}
	sQueries := make([]string, 0)
	for i, holding := range holdings {
		q := fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4)
		sQueries = append(sQueries, q)
		args = append(args, holding.Symbol, holding.Isin, holding.Quantity, holding.AvgPrice)
	}

	query += strings.Join(sQueries, ",") + ";"

	res, err = tx.Exec(query, args...)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err = res.RowsAffected()
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	if int(affected) != len(holdings) {
		tx.Rollback()
		return errors.New("not all holdings were inserted")
	}

	updateBrokerStatus = "update supported_broker set holdings_sync='t' where broker_name='zerodha';"
	res, err = tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in insertion", err.Error())
		tx.Rollback()
		return err
	}
	affected, err = res.RowsAffected()
	if err != nil {
		log.Println("Error in updation", err.Error())
		tx.Rollback()
		return err
	}
	log.Println(affected)

	return tx.Commit()
}
