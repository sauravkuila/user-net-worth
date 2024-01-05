package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/sauravkuila/portfolio-worth/external"
)

var (
	dbObj *databaseStruct
)

type databaseStruct struct {
	psql *sql.DB
}

type DatabaseInterface interface {
	GetSupportedBrokers() ([]string, error)
	// GetZerodhaHoldings() map[string]interface{}
	GetAngelOneHoldings() ([]external.HoldingsInfo, time.Time, float64, error)
	InsertAngelOneHoldings([]external.HoldingsInfo) error
}

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "portfolio"
)

func InitDb() (DatabaseInterface, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("error in db connection. error: ", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("unable to ping the db. error: ", err.Error())
		return nil, err
	}

	dbObj = &databaseStruct{
		psql: db,
	}

	return dbObj, nil
}

func CloseDb() error {
	return dbObj.psql.Close()
}
