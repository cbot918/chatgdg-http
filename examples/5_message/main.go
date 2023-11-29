package main

import (
	"chatgdg-http/examples/5_message/ws"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	port = ":8889"
)

var chat *Chat

func init() {
	chat = &Chat{
		UserCount: 0,
		Clients:   []Client{},
	}
}

func main() {
	http.HandleFunc("/ws", handleWS)
	fmt.Println(port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {

	chat.UserCount += 1

	clientID := r.URL.Query().Get("name")

	conn, rw, err := ws.Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

	client := NewClient(clientID, conn, rw)

	chanMsg, err := client.ReadMessage()
	if err != nil {
		fmt.Println("read message failed")
		return
	}
	fmt.Printf("%#+v", chanMsg)

	// write channel message
	cm := &ChanMsg{
		Chan: "userCount",
		Msg: Msg{
			Data: strconv.Itoa(chat.UserCount),
		},
	}
	err = client.WriteMessage(cm)
	if err != nil {
		fmt.Println("write message failed")
		return
	}
}
