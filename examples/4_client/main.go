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

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {

	clientID := r.URL.Query().Get("name")

	conn, rw, err := Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

	client := NewClient(clientID, conn, rw)

	inMessage, err := client.ReadMessage()
	if err != nil {
		fmt.Println("read message failed")
		return
	}
	fmt.Println(client.ID + ": " + string(inMessage))

	err = client.WriteMessage(inMessage)
	if err != nil {
		fmt.Println("write message failed")
		return
	}
}
