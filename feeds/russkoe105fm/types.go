package russkoe105fm

type feedPost struct {
	Title string `json:"title"`
	UID   string `json:"uid"`
	URL   string `json:"url"`
}

type feedResponse struct {
	Posts []feedPost `json:"posts"`
}

type postItem struct {
	Image string `json:"image"`
	Text  string `json:"text"`
}

type postResponse struct {
	Post postItem `json:"post"`
}
