package db

import (
	"database/sql"
	"log"
)

func (obj *databaseStruct) GetBrokerCred(broker string) (map[string]interface{}, error) {
	var (
		account    sql.NullString
		totpSecret sql.NullString
		userKey    sql.NullString
		passKey    sql.NullString
		appKey     sql.NullString
		secretKey  sql.NullString
		updatedOn  sql.NullTime
	)
	query := "select account, totp_secret, user_key, pass_key, app_api_key, secret_key, updated_on from creds where account=$1 limit 1;"

	row := obj.psql.QueryRow(query, broker)
	if row.Err() != nil {
		log.Println("error in querying for creds. ", row.Err().Error())
		return nil, row.Err()
	}

	err := row.Scan(&account, &totpSecret, &userKey, &passKey, &appKey, &secretKey, &updatedOn)
	if err != nil {
		log.Println("error in scanning for creds. ", err)
		return nil, err
	}

	data := make(map[string]interface{})
	data["user_key"] = userKey.String
	data["pass_key"] = passKey.String
	data["totp_secret"] = totpSecret.String
	data["app_api_key"] = appKey.String
	data["secret_key"] = secretKey.String
	data["updated_on"] = updatedOn.Time

	return data, nil
}

func (obj *databaseStruct) UpdateBrokerCred(data map[string]interface{}) error {
	return nil
}
