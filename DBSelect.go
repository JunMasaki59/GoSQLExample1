package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       string
	Created   time.Time
	Updated   time.Time
}

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(127.0.0.1:13306)/test_db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	fmt.Println("------------------")
	getRows(db)
	fmt.Println("------------------")
	getSingleRow(db, 1)
	fmt.Println("------------------")
	getSingleRow(db, 4) // 存在しないUserID
}

func getRows(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}
	defer rows.Close()
//rows.Next()で各レコードに対して操作が可能
	for rows.Next() {
		u := &User{}
//rows.Scan()で引数に渡したポインタにレコードの内容を読み込ますことができる
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Created, &u.Updated); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}
}

func getSingleRow(db *sql.DB, userID int) {
	u := &User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", userID).
		Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Created, &u.Updated)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("getSingleRow no records.")
		return
	}
	if err != nil {
		log.Fatalf("getSingleRow db.QueryRow error err:%v", err)
	}
	fmt.Println(u)
}