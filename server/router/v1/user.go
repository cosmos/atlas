package v1

// User defines the request type when updating a user record.
type User struct {
	Email string `json:"email" validate:"required,email"`
}
