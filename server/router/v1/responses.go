package v1

// ModuleStars defines the HTTP response type for the total nubmer of favorites
// for a module.
type ModuleStars struct {
	Stars int64 `json:"stars"`
}
