package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isBetween(x int, min int, max int) bool {
	return x >= min || x < max

}

func showResponse(resp *http.Response) {
	if isBetween(resp.StatusCode, 200, 300) {
		bytes, err := ioutil.ReadAll(resp.Body)
		checkError(err)

		fmt.Println(string(bytes), "status= ", resp.StatusCode)
	} else {
		fmt.Println("failed for ", resp.StatusCode, resp.Header)
	}
}

func get(url string) {

	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	showResponse(resp)
}

type Post struct {
	Text string `json:"text"`
}

func (post Post) ToString() string {
	return "{\"text\": \"" + post.Text + "\"}"
}

func sendRequest(url string, jsonStr string, method string) {

	jsonBytes := []byte(jsonStr)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	checkError(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	showResponse(resp)

	decoder := json.NewDecoder(resp.Body)

	item := &Post{}
	decoder.Decode(item)
	//fmt.Println(item.ToString())

}

var home = "http://localhost:8080/"
var url = "http://localhost:8080/api/v1/text"


func main() {
	//get(home)
	get(url)
    n:= 100
	for i:=0; i<n; i++ {
		text:= "{\"text\": \"post"+strconv.Itoa(i)+"\"}"
		sendRequest(url, string(text), "POST")
	}
	for i:=0; i<n/2; i++ {
		text:= "{\"text\": \"post-updated"+strconv.Itoa(i)+"\"}"
		sendRequest(url+"?id="+strconv.Itoa(i), string(text), "PUT")
	}
	for i:=n/8; i<n/4; i++ {
		sendRequest(url+"?id="+strconv.Itoa(i), "", "DELETE")
	}

	get(url)

}
