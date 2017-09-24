package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AppRootHandler returns the index.html when '/' is hit
func AppRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("index.html")

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func main() {
	http.HandleFunc("/", AppRootHandler)
	log.Fatal(http.ListenAndServe(":1337", nil))
}
