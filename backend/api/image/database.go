package image

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

var imgOutDir = "./../../../data/database/images"
var Posts Database = make(Database)

func fetchPost(post *Post) []byte {

	path := post.Path
	if file, err := ioutil.ReadFile(imgOutDir + "/" + path); err != nil {
		var file []byte
		return file
	} else {
		return file
	}
}

func deleteItem(id int) bool {
	post := findPost(id)
	if err := os.Remove(post.Path); err != nil {
		return false
	}
	delete(Posts, id)
	return true
}

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
		return &Post{"err: index " + strconv.Itoa(id) + " Not Found", ""}
	}
}

func addPost(post *Post, imgBytes multipart.File) bool {
	return addPostById(post, len(Posts), imgBytes)
}

func store(path string, imgBytes multipart.File) bool {

	if temp, err := os.Create(imgOutDir + "/" + path); err == nil {
		defer temp.Close()
		if _, err := io.Copy(temp, imgBytes); err == nil {
			return true
		}
	}
	return false
}

func addPostById(post *Post, id int, imgBytes multipart.File) bool {
	store(post.Path, imgBytes)
	Posts[id] = post
	return true
}

func updatePost(post *Post, id int, imgBytes multipart.File) bool {
	if !idIsOk(id) {
		return false
	}
	deleteItem(id)
	return addPostById(post, id, imgBytes)
}

func removePost(id int) bool {

	if !idIsOk(id) {
		return false
	}
	return deleteItem(id)
}
