package iestafeta

type post struct {
	Description string `xml:"description"`
	Content     string `xml:"content:encoded:"`
	GUID        string `xml:"guid"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Title       string `xml:"title"`
}

type channel struct {
	Items []post `xml:"item"`
}

type data struct {
	Channel channel `xml:"channel"`
}
