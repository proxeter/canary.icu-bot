package iestafeta

type source struct {
	Data string `xml:",chardata"`
}

type post struct {
	Content source `xml:"encoded"`
	GUID    string `xml:"guid"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Title   string `xml:"title"`
}

type channel struct {
	Items []post `xml:"item"`
}

type data struct {
	Channel channel `xml:"channel"`
}
