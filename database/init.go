package database

import (
	"container/list"
	"database/sql"
	"fmt"
	"log"
)

// CachedTableMeta provides cached data of the DB
var CachedTableMeta map[string]*list.List

// Db is an instance of db connection
var Db *sql.DB

// InitService initializes the service
type InitService struct {
	Db *sql.DB
}

// Init puts table and column names in memory
func (s InitService) Init() {
	Db = s.Db
	rows, err := Db.Query("SELECT table_name as show_tables FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema');")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	cache := make(map[string]*list.List)
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		cache[tableName] = list.New()
		columnQuery := fmt.Sprintf("SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '%s';", tableName)
		cols, err := s.Db.Query(columnQuery)
		for cols.Next() {
			var colName string
			err := cols.Scan(&colName)
			if err != nil {
				log.Fatal(err)
			}
			cache[tableName].PushBack(colName)
		}
		defer cols.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	CachedTableMeta = cache
}
