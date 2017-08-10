package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("./"))))
	//log.Fatal(http.ListenAndServeTLS(":443", http.FileServer(http.Dir("./"))))
}
