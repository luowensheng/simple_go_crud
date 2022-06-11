package main

type Post struct {
	Text string `json:"text"`
}

func (post Post) ToString() string {
	return "{\"text\": \"" + post.Text + "\"}"
}
