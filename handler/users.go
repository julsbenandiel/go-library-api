package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julsbenandiel/go-library-api/database"
)

type BirthDatePayload struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
}

type TempCreatePayload struct {
	*database.CreateUserParams
	BirthDate BirthDatePayload `json:"birth_date"`
}

func (api *User) HandleGetUsers(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	users, err := api.Queries.GetUsers(ctx)

	log.Print(users)

	if err != nil {
		log.Fatal("Cannot get users. Err ", err)
	}

	// var tempDataSlice []map[string]interface{}

	// for _, user := range users {
	// 	d, err := time.Parse(time.RFC3339, user.BirthDate)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	tempDataSlice = append(tempDataSlice, map[string]interface{}{
	// 		"tanga":      "ka bobo",
	// 		"ID":         user.ID,
	// 		"FirstName":  user.FirstName,
	// 		"LastName":   user.LastName,
	// 		"Email":      user.Email,
	// 		"Username":   user.Username,
	// 		"Address":    user.Address,
	// 		"CreatedAt":  user.CreatedAt,
	// 		"UpdpatedAt": user.UpdpatedAt,
	// 		"BirthDate":  d,
	// 	})
	// }

	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func (api *User) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	payload := database.CreateUserPayload{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		log.Fatal("Cannot decode payload: ", err)
		return
	}

	data := database.CreateUserParams{
		ID:         uuid.New(),
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Email:      payload.Email,
		Username:   payload.Username,
		Address:    payload.Address,
		CreatedAt:  time.Now(),
		UpdpatedAt: time.Now(),
	}

	parsedDate, err := time.Parse("02-01-2006", payload.BirthDate)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	data.BirthDate = parsedDate

	createdUser, err := api.Queries.CreateUser(r.Context(), data)
	if err != nil {
		log.Print(createdUser.ID)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("User Created"))
}
