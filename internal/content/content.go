package content

// Sourcer defines common behaviour of content sourcers.
type Sourcer interface {
	Handle() (interface{}, error)
}
