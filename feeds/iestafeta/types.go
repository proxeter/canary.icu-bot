package iestafeta

type post struct {
	Link    string `xml:"link"`
	Message string `xml:"description"`
	Title   string `xml:"title"`
	UID     string `xml:"guid"`
}

type channel struct {
	Items []post `xml:"item"`
}

type data struct {
	Channel channel `xml:"channel"`
}
