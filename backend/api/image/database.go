package image

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

var DIR, _ = os.Getwd()
var imgOutDir = DIR + "/../data/database/images"
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

func createFilePath(id int) string {
	return strconv.Itoa(id) + ".png"
}

func addPost(imgBytes multipart.File) bool {
	var id = len(Posts)
	var path = createFilePath(id)
	var post = NewPost(path, path)
	return addPostById(post, id, imgBytes)
}

func store(path string, imgBytes multipart.File) bool {

	if temp, err := os.Create(imgOutDir + "/" + path); err == nil {

		defer temp.Close()
		if _, err := io.Copy(temp, imgBytes); err == nil {
			return true
		} else {
			fmt.Println(err)
			return false
		}
	} else {
		fmt.Println(err)
		return false
	}
}

func addPostById(post *Post, id int, imgBytes multipart.File) bool {
	success := store(post.Path, imgBytes)
	if success {
		Posts[id] = post
	}
	return success
}

func updatePost(id int, imgBytes multipart.File) bool {
	if !idIsOk(id) {
		return false
	}
	deleteItem(id)
	path := createFilePath(id)
	post := NewPost(path, path)
	return addPostById(post, id, imgBytes)
}

func removePost(id int) bool {

	if !idIsOk(id) {
		return false
	}
	return deleteItem(id)
}
