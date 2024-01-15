package broker

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *brokerDbSt) InsertAngelOneHoldings(holdings []external.HoldingsInfo) error {
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
	updateBrokerStatus := "update supported_sources set holdings_sync='t' where source_name='angelone';"
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

	updateBrokerStatus = "update supported_sources set holdings_sync='t' where source_name='angelone';"
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

func (obj *brokerDbSt) InsertIDirectHoldings(holdings []external.HoldingsInfo) error {
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

	updateBrokerStatus := "update supported_sources set holdings_sync='f' where source_name='idirect';"
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

	updateBrokerStatus = "update supported_sources set holdings_sync='t' where source_name='idirect';"
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

func (obj *brokerDbSt) InsertZerodhaHoldings(holdings []external.HoldingsInfo) error {
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

	updateBrokerStatus := "update supported_sources set holdings_sync='f' where source_name='zerodha';"
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

	updateBrokerStatus = "update supported_sources set holdings_sync='t' where source_name='zerodha';"
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
