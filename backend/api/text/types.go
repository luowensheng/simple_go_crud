package text

type Result struct {
	Response string `json:"response"`
	Method   string `json:"method"`
	Success  bool   `json:"success"`
	Post     *Post  `json:"post"`
}

type Results struct {
	Response string   `json:"response"`
	Method   string   `json:"method"`
	Success  bool     `json:"success"`
	Database Database `json:"database"`
}

type Post struct {
	Text string `json:"text"`
}

type Database = map[int]*Post
