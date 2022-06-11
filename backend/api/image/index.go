package image

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handleCreate(rw http.ResponseWriter, r *http.Request) {

	if file, _, err := r.FormFile("image"); err == nil {

		if success := addPost(file); success {
			fmt.Println("added success")
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusCreated)
			json.NewEncoder(rw).Encode(&Result{"", "POST", success, nil})

		} else {
			fmt.Println("added failed")

			rw.WriteHeader(http.StatusNoContent)
			json.NewEncoder(rw).Encode(&Result{"", "POST", success, nil})
		}
		fmt.Println("DONE!")
		return

	} else {
		fmt.Println("reading fail")

		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(&Result{"fail", "POST", false, nil})
	}

}

func handleRead(rw http.ResponseWriter, r *http.Request) {
	id := parseIdFromUrl(r)
	post := findPost(id)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	if getIdFromUrl(r) == "" {
		json.NewEncoder(rw).Encode(&Results{"", "GET", true, getAllPosts()})
	} else {
		json.NewEncoder(rw).Encode(&Result{"", "GET", true, post})
	}
}

func handleUpdate(rw http.ResponseWriter, r *http.Request) {
	post := &Post{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(post)
	id := parseIdFromUrl(r)
	file, _, err := r.FormFile("image")
	if err != nil {
		success := updatePost(id, file)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(&Result{"", "PUT", success, findPost(id)})
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}

}
func handleDelete(rw http.ResponseWriter, r *http.Request) {

	id := parseIdFromUrl(r)
	success := removePost(id)
	rw.Header().Set("Content-Type", "application/json")
	if success {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(&Result{"", "DELETE", success, findPost(id)})
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(&Result{"", "DELETE", success, findPost(id)})
	}

}

func getIdFromUrl(r *http.Request) string {
	return r.URL.Query().Get("id")
}
func parseIdFromUrl(r *http.Request) int {
	if i, err := strconv.Atoi(getIdFromUrl(r)); err == nil {
		return i
	} else {
		return -1
	}

}

func ImageHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request")
	fmt.Println(parseIdFromUrl(r))
	// file, h, err := r.FormFile("image")
	// fmt.Println(file, h, err)

	fmt.Println(r.Method)

	switch r.Method {
	case "GET":
		handleRead(rw, r)
	case "POST":
		handleCreate(rw, r)
	case "PUT":
		handleUpdate(rw, r)
	case "DELETE":
		handleDelete(rw, r)
	default:
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Invalid Method ("+r.Method+")")
	}
	fmt.Println("handled request")

}
