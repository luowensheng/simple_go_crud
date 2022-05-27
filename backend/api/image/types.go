package image

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
	Path    string `json:"path"`
	OutPath string `json:"-"`
}

type Database = map[int]*Post
