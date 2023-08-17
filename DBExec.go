package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(127.0.0.1:13306)/test_db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	userID := insertUser(db, "Satoshi", "Yamada", 27)
	insertPost(db, userID, "hello world")
}

func insertUser(db *sql.DB, firstName, lastName string, age int) int64 {
	res, err := db.Exec(
		"INSERT INTO users (first_name, last_name, age) VALUES (?, ?, ?)",
		firstName,
		lastName,
		age,
	)
	if err != nil {
		log.Fatalf("insertUser db.Exec error err:%v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("insertUser res.LastInsertId error err:%v", err)
	}
	return id
}

func insertPost(db *sql.DB, userID int64, content string) int64 {
	res, err := db.Exec("INSERT INTO posts (user_id, content) VALUES (?, ?)",
		userID,
		content,
	)
	if err != nil {
		log.Fatalf("insertPost db.Exec error err:%v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("insertPost res.LastInsertId error err:%v", err)
	}
	return id
}