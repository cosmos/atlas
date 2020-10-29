package v1

// ModuleStars defines the HTTP response type for the total nubmer of favorites
// for a module.
type ModuleStars struct {
	Total int64 `json:"total"`
}
