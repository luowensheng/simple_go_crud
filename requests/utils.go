package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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



func excecuteRequest(req *http.Request) *http.Response {
	fmt.Println("Sending request")
	client := &http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	fmt.Println("Received Response")

	// defer resp.Body.Close()
	return resp
}
func sendRequest(url string, jsonStr string, method string) {

	jsonBytes := []byte(jsonStr)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	checkError(err)

	req.Header.Set("Content-Type", "application/json")
	resp := excecuteRequest(req)
	showResponse(resp)

	decoder := json.NewDecoder(resp.Body)

	item := &Post{}
	decoder.Decode(item)
	//fmt.Println(item.ToString())

}



func sendFileData(url, path, fieldname, filename string) {

	file, err := os.Open(path)
	checkError(err)

	pr, pw := io.Pipe()
	defer pr.Close()

	form := multipart.NewWriter(pw)

	onerror := func(err error) {
		if err != nil {
			pw.CloseWithError(err)
			checkError(err)
		}
	}
	go func() {
		w, err := form.CreateFormFile(fieldname, filename)

		onerror(err)
		_, err = io.Copy(w, file)
		onerror(err)

		form.Close()
		pw.Close()

	}()

	req, err := http.NewRequest(http.MethodPost, url, pr)
	checkError(err)
	req.Header.Set("Content-Type", form.FormDataContentType())
	resp := excecuteRequest(req)

	showResponse(resp)

	resp.Body.Close()

}

