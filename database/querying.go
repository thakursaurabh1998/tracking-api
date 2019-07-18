package database

import (
	"database/sql"
	"fmt"
)

type Data struct {
	A    int
	B, C float64
}

func DBQuery(data Data, tableName string) (sql.Result, error) {
	ok, err := Db.Exec(fmt.Sprintf("INSERT INTO %s (events_master_fk, buy_amount, sell_amount) VALUES ($1, $2, $3);", tableName), data.A, data.B, data.C)
	// defer Db.Close()
	return ok, err
}
