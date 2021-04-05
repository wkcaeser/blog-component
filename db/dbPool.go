package db

import (
	"blog-component/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DbPool *sql.DB

func Stmt(sql string) *sql.Stmt {
	stmt, err := DbPool.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	return stmt
}

func init() {

	db, err := sql.Open("mysql", config.GlobalConf.Mysql.Url)
	if err != nil {
		log.Panicln("db err: ", err.Error())
	}
	db.SetMaxOpenConns(config.GlobalConf.Mysql.MaxOpenConns)
	db.SetMaxIdleConns(config.GlobalConf.Mysql.MaxIdleConns)
	DbPool = db
}
