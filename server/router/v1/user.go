package v1

// User defines the request type when updating a user record.
type User struct {
	Email string `json:"email" validate:"required,email"`
}

// Token defines the request type when creating a new user API token.
type Token struct {
	Name string `json:"name" validate:"required"`
}
