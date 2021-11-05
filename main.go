package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Sockets!")
	// fmt.Println("Endpoint Hit")
	// ws, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// go echo(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/socket", wsEndpoint)
}

func main() {
	fmt.Println("Hello, World!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
