package v1

// User defines the request type when updating a user record.
type User struct {
	Email string `json:"email" validate:"required,email"`
}

// Token defines the request type when creating a new user API token.
type Token struct {
	Name string `json:"name" validate:"required"`
}

// ModuleInvite defines the request type when inviting a user as an owner to a
// module.
type ModuleInvite struct {
	ModuleID uint   `json:"module_id" validate:"required,gte=1"`
	User     string `json:"user" validate:"required"`
}
