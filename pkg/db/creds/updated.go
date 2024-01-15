package creds

import (
	"fmt"
	"log"
	"strings"
)

func (obj *credDbSt) UpdateBrokerCred(data map[string]interface{}) error {
	query := "update creds set "

	i := 1
	args := make([]interface{}, 0)
	q := make([]string, 0)
	var accountValue interface{}
	for k, v := range data {
		if k == "account" {
			accountValue = v
			continue
		}
		q = append(q, fmt.Sprintf("%s=$%d", k, i))
		args = append(args, v)
		i += 1
	}

	query += strings.Join(q, ",")
	query += fmt.Sprintf(" where account = $%d;", i)
	args = append(args, accountValue)

	res, err := obj.psql.Exec(query, args...)
	if err != nil {
		log.Println("unable to update the creds. ", err.Error())
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("unable to fetch affected rows. ", err.Error())
		return err
	}
	if affected != 1 {
		return fmt.Errorf("not updated")
	}
	return nil
}
