package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port               = ":8889"
	magicWebSocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

func main() {
	http.HandleFunc("/ws", handleWS)
	fmt.Println(port)
	http.ListenAndServe(port, nil)
}

func handleWS(w http.ResponseWriter, r *http.Request) {

	_, _, err := Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

}
