package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func GetMySqlDatabase() *bun.DB {
	dsn := "root@tcp(127.0.0.1:3306)/token"
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	mySqlDB := bun.NewDB(sqldb, mysqldialect.New())

	if err := mySqlDB.Ping(); err != nil {
		log.Println(err)
	}
	return mySqlDB
}
