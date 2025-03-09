package main

import (
	"fmt"
	"context"
	"os"
	"github.com/jackc/pgx/v5"
)


func dbCon () {
	cfg, err := LoadConfig()

	if err != nil {
		fmt.Fprintf(os.Stderr, "config failed: %v\n", err)
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), cfg.DB.URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	fmt.Println(name, weight)
}