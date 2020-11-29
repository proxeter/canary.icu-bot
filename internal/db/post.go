package db

// Make - fill post with factory values
func (p *Post) Make(uid string, f Factory, bucket []byte) (bool, error) {
	r, err := f(uid)

	if err != nil {
		return false, err
	}

	d, isNew, err := GetPersistentPost(uid, r)

	p.Link = d.Link
	p.Message = d.Message
	p.PreviewImage = d.PreviewImage
	p.Title = d.Title
	p.UID = d.UID

	return isNew, nil
}
