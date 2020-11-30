package db

// Factory - function for getting post metadata
type Factory func(id string) (Post, error)

// Post db.Post model
type Post struct {
	ID           string
	Link         string
	Message      string
	PreviewImage string
	Timestamp    int64
	Title        string
}
