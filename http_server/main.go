package main

import (
	"io"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello,world\n")
}

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Home Page\n")
}

func main() {
	http.HandleFunc("/hello", helloServer)
	http.HandleFunc("/", home)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
