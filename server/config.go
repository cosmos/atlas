package server

// Config defines a configuration abstraction provided to a Server type in order
// to be instantiated.
type Config interface {
	String(path string) string
	Int(path string) int
	Ints(path string) []int
}
