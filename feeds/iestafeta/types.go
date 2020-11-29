package iestafeta

type post struct {
	Description string `xml:"description"`
	GUID        string `xml:"guid"`
	Link        string `xml:"link"`
	Title       string `xml:"title"`
}

type channel struct {
	Items []post `xml:"item"`
}

type data struct {
	Channel channel `xml:"channel"`
}
