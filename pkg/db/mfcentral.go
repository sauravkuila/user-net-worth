package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *databaseStruct) InsertMfCentralHoldings(holdings []external.MfHoldingsInfo) error {
	//flush db
	tx, err := obj.psql.Begin()
	if err != nil {
		log.Println("txn error. ", err.Error())
		return err
	}
	delQuery := "delete from mf_central;"
	_, err = tx.Exec(delQuery)
	if err != nil {
		log.Println("failed to flush. ", err.Error())
		tx.Rollback()
		return err
	}

	updateBrokerStatus := "update supported_broker set holdings_sync='f' where broker_name='mfcentral';"
	res, err := tx.Exec(updateBrokerStatus)
	if err != nil {
		log.Println("Error in updation", err.Error())
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

	query := "insert into mf_central(folio, scheme_name, isin, quantity, price, cost_price, curr_price) values "
	var args []interface{}
	sQueries := make([]string, 0)
	for i, holding := range holdings {
		q := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7)
		sQueries = append(sQueries, q)
		args = append(args, holding.Folio, holding.Name, holding.Isin, holding.Quantity, holding.AvgPrice, holding.InvestedVal, holding.CurrentVal)
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

	updateBrokerStatus = "update supported_broker set holdings_sync='t' where broker_name='mfcentral';"
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

func (obj *databaseStruct) GetMfCentralHoldings() ([]external.MfHoldingsInfo, float64, error) {
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
