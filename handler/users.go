package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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

	payload := &TempCreatePayload{
		BirthDate: BirthDatePayload{},
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		log.Fatal("Cannot decode payload: ", err)
		return
	}

	birthDate := time.Date(
		payload.BirthDate.Year,
		time.Month(payload.BirthDate.Month),
		payload.BirthDate.Date,
		0, 0, 0, 0, time.Local,
	)

	log.Println("now:", time.Now())
	log.Println("birth date:", birthDate)

	createdUser, err := api.Queries.CreateUser(r.Context(), database.CreateUserParams{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Username:  payload.Username,
		Address:   payload.Address,
		BirthDate: pgtype.Date{Time: birthDate, InfinityModifier: pgtype.Finite},
	})

	cu, _ := json.MarshalIndent(createdUser, "", " ")
	log.Println(string(cu))

	if err != nil {
		log.Fatal("Failed to create user.", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("User Created"))
}
