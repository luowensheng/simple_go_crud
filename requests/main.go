package main

import "strconv"

// "sync"

var home = "http://localhost:8080/"
var url = "http://localhost:8080/api/v1/text"
var img_url = "http://localhost:8080/api/v1/image"

func text() {
	//get(home)
	get(url)
	n := 100
	for i := 0; i < n; i++ {
		text := "{\"text\": \"post" + strconv.Itoa(i) + "\"}"
		sendRequest(url, string(text), "POST")
	}
	for i := 0; i < n/2; i++ {
		text := "{\"text\": \"post-updated" + strconv.Itoa(i) + "\"}"
		sendRequest(url+"?id="+strconv.Itoa(i), string(text), "PUT")
	}
	for i := n / 8; i < n/4; i++ {
		sendRequest(url+"?id="+strconv.Itoa(i), "", "DELETE")
	}

	get(url)

}

func image(){
	path := "/home/oliver/Documents/coding/go/netstuff/requests/data/car.jpeg"
	fieldnmae := "image"
	filename := "car.jpeg"
	sendFileData(img_url, path, fieldnmae, filename)
}
func main() {
     
	image()

	//text()
}
