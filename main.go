package main

import (
	"HFT/ftx_ws"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to DB
	log.Println("Connecting to db...")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	log.Println("Done!")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	ftx_ws.SubscribeOB(conn)
}
