package creds

import "database/sql"

type credDbSt struct {
	psql *sql.DB
}

type CredsDatabaseInterface interface {
	//return  map of source type vs sources against it
	GetSupportedSources() (map[string][]string, error)
	//get the broker specific cred
	GetBrokerCred(broker string) (map[string]interface{}, error)
	//update the broker creds
	UpdateBrokerCred(data map[string]interface{}) error
}

func NewCredsDbInterface(db *sql.DB) CredsDatabaseInterface {
	return &credDbSt{
		psql: db,
	}
}
