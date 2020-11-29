package db

// Factory - function for getting post metadata
type Factory func(id string) (Post, error)

// Post db.Post model
type Post struct {
	Link         string
	Message      string
	PreviewImage string
	Title        string
	ID           string
}
