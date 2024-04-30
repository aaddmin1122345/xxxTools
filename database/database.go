package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Db *sql.DB

func dbInit() {
	db, err := sql.Open("sqlite3", "poc.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {

		}
	}(db)
}

func QueryPoc(db *sql.DB, pocName string) (string, error) {
	dbInit()
	var (
		payload    string
		httpMethod string
		text       string
	)

	// 执行查询语句
	row := db.QueryRow("SELECT httpMethod, other FROM poc WHERE pocName = ?", pocName)

	// 将查询结果赋值给变量
	err := row.Scan(&httpMethod, &text)
	if err != nil {
		return "查询有误", err
	}

	// 返回查询结果
	return payload, nil
}
