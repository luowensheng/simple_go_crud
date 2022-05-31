# simple_go_crud

```go
func handleCreate(rw http.ResponseWriter, r *http.Request) {
	post := &Post{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(post)

	success := addPost(post)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&Result{"", "POST", success, post})
}

func handleRead(rw http.ResponseWriter, r *http.Request) {
	id := parseIdFromUrl(r)
	post := findPost(id)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	if getIdFromUrl(r) == "" {
		json.NewEncoder(rw).Encode(&Results{"", "GET", true, getAllPosts()})
		return
	}
	json.NewEncoder(rw).Encode(&Result{"", "GET", true, post})
}

func handleUpdate(rw http.ResponseWriter, r *http.Request) {
	post := &Post{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(post)
	id := parseIdFromUrl(r)
	success := updatePost(post, id)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&Result{"", "PUT", success, findPost(id)})
}
func handleDelete(rw http.ResponseWriter, r *http.Request) {

	id := parseIdFromUrl(r)
	success := removePost(id)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&Result{"", "DELETE", success, findPost(id)})
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

func TextHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Println(parseIdFromUrl(r))

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
}

```
