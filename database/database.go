package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Db *sql.DB

func dbInit() {
	var err error
	Db, err = sql.Open("sqlite3", "database/poc.db")
	if err != nil {
		log.Fatal(err)
	}

}

func QueryPoc(pocName string) (string, error) {
	if Db == nil {
		dbInit()
	}
	var (
		payload    string
		httpMethod string
		text       string
	)

	// 执行查询语句
	row := Db.QueryRow("SELECT payload, httpMethod, other FROM poc WHERE pocName = ?", pocName)

	// 将查询结果赋值给变量
	err := row.Scan(&payload, &httpMethod, &text)
	if err != nil {
		return "查询有误", err
	}

	// 返回查询结果
	return payload, nil
}
