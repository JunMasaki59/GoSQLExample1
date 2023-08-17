package main

import (
	"database/sql"
	"fmt"
	"log"
	//"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(127.0.0.1:13306)/test_db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	fmt.Printf("%+v\n", db.Stats())
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	fmt.Printf("%+v\n", db.Stats())

	var wg sync.WaitGroup
	s := time.Now()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go request(&wg, db, i)
	}
	wg.Wait()
	e := time.Now()
	fmt.Printf("処理秒数: %v\n", e.Sub(s).Round(time.Millisecond))
}

func request(wg *sync.WaitGroup, db *sql.DB, i int) {
	defer wg.Done()

	fmt.Printf("[request start] i: %v\n", i)
	defer fmt.Printf("[request end] i: %v\n", i)

	rows, err := db.Query("SELECT * FROM tests")
	if err != nil {
		log.Fatalf("request db.Query error err:%v", err)
	}
	defer rows.Close()
}