package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/sauravkuila/portfolio-worth/pkg/db/broker"
	"github.com/sauravkuila/portfolio-worth/pkg/db/creds"
	"github.com/sauravkuila/portfolio-worth/pkg/db/mutualfund"
)

var (
	dbObj *databaseStruct
)

type databaseStruct struct {
	psql *sql.DB
	broker.BrokerDatabaseInterface
	mutualfund.MutualFundDatabaseInterface
	creds.CredsDatabaseInterface
}

type DatabaseInterface interface {
	broker.BrokerDatabaseInterface
	mutualfund.MutualFundDatabaseInterface
	creds.CredsDatabaseInterface
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
		db,
		broker.NewBrokerDBInterface(db),
		mutualfund.NewMutualfundInterfaceObj(db),
		creds.NewCredsDbInterface(db),
	}

	return dbObj, nil
}

func CloseDb() error {
	return dbObj.psql.Close()
}
