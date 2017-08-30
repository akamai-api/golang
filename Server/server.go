package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK! I got you Bro! -> ")
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":9143", nil)
}
