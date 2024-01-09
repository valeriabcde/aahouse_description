package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	urlExample := "postgres://house:123@localhost:5436/test_db"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("ready")
	}

	ShowDb()

	//defer conn.Close(context.Background())
	//
	//nameInsert := "Вова"
	////var name string
	////var id string
	//tx, err := conn.Begin(context.Background())
	//if err != nil {
	//  fmt.Println(err)
	//}
	//defer tx.Rollback(context.Background())
	//tx.Exec(context.Background(), "insert into users (name) values ($1)", nameInsert)
	//tx.Commit(context.Background())
	//rows, err := conn.Query(context.Background(), "select * from users")
	//if err != nil {
	//  fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//  os.Exit(1)
	//}
	//for rows.Next() {
	//  rows.Scan(&name, &id)
	//  fmt.Println("name - ", name, "id - ", id)
	//}
}
