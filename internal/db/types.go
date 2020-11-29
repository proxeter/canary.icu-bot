package db

// Factory - function for getting post metadata
type Factory func(uid string) (Post, error)

// Post db.Post model
type Post struct {
	Link         string
	Message      string
	PreviewImage string
	Title        string
	UID          string
}
