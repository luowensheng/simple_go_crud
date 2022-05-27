package text

import "strconv"


var Posts Database = make(Database)

func idIsOk(id int) bool {
	return Posts[id] != nil

}
func getAllPosts() Database {
	return Posts
}
func findPost(id int) *Post {

	if idIsOk(id) {
		return Posts[id]
	} else {
		return &Post{"err: index " + strconv.Itoa(id) + " Not Found"}
	}
}

func addPost(post *Post) bool {
	Posts[len(Posts)] = post
	return true

}
func updatePost(post *Post, id int) bool {
	if !idIsOk(id) {
		return false
	}
	Posts[id] = post
	return true
}

func removePost(id int) bool {
	if !idIsOk(id) {
		return false
	}
	delete(Posts, id)
	return true
}
