package creds

import (
	"database/sql"
	"log"
)

func (obj *credDbSt) GetBrokerCred(broker string) (map[string]interface{}, error) {
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
	data["account"] = account.String
	data["user_key"] = userKey.String
	data["pass_key"] = passKey.String
	data["totp_secret"] = totpSecret.String
	data["app_api_key"] = appKey.String
	data["secret_key"] = secretKey.String
	data["updated_on"] = updatedOn.Time

	return data, nil
}

func (obj *credDbSt) GetSupportedSources() (map[string][]string, error) {
	query := "select source_name, source_type from supported_sources group by source_type,source_name;"
	rows, err := obj.psql.Query(query)
	if err != nil {
		return nil, err
	}
	resp := make(map[string][]string)
	for rows.Next() {
		var (
			sourceName sql.NullString
			sourceType sql.NullString
		)
		err := rows.Scan(&sourceName, &sourceType)
		if err != nil {
			return nil, err
		}
		resp[sourceType.String] = append(resp[sourceType.String], sourceName.String)
	}
	return resp, nil
}
