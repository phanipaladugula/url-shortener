package main

import(
	"fmt"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func initDB(){
	connStr := os.Getenv("DB_URL")
    if connStr == "" {
        connStr = "postgres://postgres:Phani@123@localhost:5432/postgres"
    }
	var err error
	dbPool, err = pgxpool.New(context.Background(),connStr)

	if err != nil{
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v \n",err)
		os.Exit(1)
	}
	fmt.Println("Connected to Postgres SQL successfully!")
}