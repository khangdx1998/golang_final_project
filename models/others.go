package  models

type Condition struct {
	Field string
	Value string
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}