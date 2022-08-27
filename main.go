package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./build/html"))
	router.Handle("/", fs)
	fmt.Println("http://localhost:5000")
	http.ListenAndServe(":5000", router)
}
