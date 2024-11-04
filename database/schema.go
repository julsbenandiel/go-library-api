package database

type CreateUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
}
