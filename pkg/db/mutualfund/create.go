package mutualfund

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *mutualfundDbSt) InsertMfCentralHoldings(holdings []external.MfHoldingsInfo) error {
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

	updateBrokerStatus := "update supported_sources set holdings_sync='f' where source_name='mfcentral';"
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

	updateBrokerStatus = "update supported_sources set holdings_sync='t' where source_name='mfcentral';"
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
