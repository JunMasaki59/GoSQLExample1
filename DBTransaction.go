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

	transaction(db)
}

func transaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if err := tx.Rollback(); err != nil {
				log.Fatalf("transaction rollback error err:%v", err)
			}
		}
	}()

	userID, err := insertUserTx(tx, "Satoshi", "Yamada", 27)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("transaction rollback error err:%v", err)
		}
		log.Fatalf("transaction insertUserTx error err:%v", err)
	}

	_, err = insertPostTx(tx, *userID, "hello world")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("transaction rollback error err:%v", err)
		}
		log.Fatalf("transaction insertPostTx error err:%v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("transaction commit error err:%v", err)
	}
}

func insertUserTx(tx *sql.Tx, firstName, lastName string, age int) (*int64, error) {
	res, err := tx.Exec(
		"INSERT INTO users (first_name, last_name, age) VALUES (?, ?, ?)",
		firstName,
		lastName,
		age,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func insertPostTx(tx *sql.Tx, userID int64, content string) (*int64, error) {
	res, err := tx.Exec("INSERT INTO posts (user_id, content) VALUES (?, ?)",
		userID,
		content,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}