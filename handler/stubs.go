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

func (api *API) GetStubsWithUser(w http.ResponseWriter, _ *http.Request) {

	stubs, err := api.Queries.GetStubs(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(stubs)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func (api *API) GetStubsByUser(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("id")

	stubs, err := api.Queries.GetUserStubs(r.Context(), uuid.MustParse(userId))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(stubs)

	data, err := json.Marshal(stubs)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func (api *API) CreateStub(w http.ResponseWriter, r *http.Request) {

	payload := database.CreateStubParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdpatedAt: time.Now(),
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		log.Fatal(err)
	}

	data, err := api.Queries.CreateStub(r.Context(), payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(data.ID.String()))
}
