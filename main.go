package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/julsbenandiel/go-library-api/database"
	"github.com/julsbenandiel/go-library-api/handler"
)

func main() {
	loadEnvVars()

	db := connectToDB()
	queries := database.New(db)

	router := http.NewServeMux()

	userHandler := &handler.User{Queries: queries}

	sampleDate := "07-09-1996"
	parsedDate, err := time.Parse("02-01-2006", sampleDate)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsedDate)

	router.HandleFunc("GET /user", userHandler.HandleGetUsers)
	router.HandleFunc("POST /user", userHandler.HandleCreateUser)

	// apiv1 := http.NewServeMux()
	// apiv1.Handle("/api/", http.StripPrefix("/api", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started on port 8080")
	server.ListenAndServe()
}

func connectToDB() *pgx.Conn {
	ctx := context.Background()
	dbConnStr := os.Getenv("POSTGRES_CONNECTING_STRING")

	db, err := pgx.Connect(ctx, dbConnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatal("Cannot ping db: ", err)
	}

	log.Println("-> connected to db")

	return db
}

func loadEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load environment variables.")
	}

	fmt.Println("-> .env successfully loaded")
}
