package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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

	// queries := database.New(db)

	// user, err := queries.CreateUser(ctx, database.CreateUserParams{
	// 	FirstName: pgtype.Text{String: "John"},
	// 	LastName:  pgtype.Text{String: "Doe"},
	// 	Email:     "julsbenandiel@gmail.com",
	// 	Username:  "julsbenandiel",
	// 	Address:   "Bijlmerplein 8888",
	// 	BirthDate: pgtype.Date{Time: time.Date(1996, time.September, 7, 0, 0, 0, 0, time.UTC)},
	// })
	// if err != nil {
	// 	log.Fatal("cannot create User: ", err)
	// }

	// log.Println("User created ", user)
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
