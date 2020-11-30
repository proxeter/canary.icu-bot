package db

// Make - fill post with factory values
func (p *Post) Make(id string, f Factory, bucket []byte) (bool, error) {
	r, err := f(id)

	if err != nil {
		return false, err
	}

	d, isNew, err := GetPersistentPost(r)

	p.ID = d.ID
	p.Link = d.Link
	p.Message = d.Message
	p.PreviewImage = d.PreviewImage
	p.Timestamp = d.Timestamp
	p.Title = d.Title

	return isNew, nil
}
