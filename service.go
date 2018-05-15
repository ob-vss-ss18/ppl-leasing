package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, Dev Heroku!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)                     // set router
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
