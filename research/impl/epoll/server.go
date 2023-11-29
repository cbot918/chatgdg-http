package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/websocket"
)

const port = ":8001"

var epoller *epoll

func main() {

	// Start epoll
	var err error
	epoller, err = MkEpoll()
	if err != nil {
		panic(err)
	}

	go Start()

	http.HandleFunc("/ws", wsHandler)

	fmt.Println(port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	fmt.Println("in")
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://127.0.0.1:5500"
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := epoller.Add(conn); err != nil {
		log.Printf("Failed to add connection")
		conn.Close()
	}
}

func Start() {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			_, msg, err := conn.ReadMessage()
			if err != nil {
				if err := epoller.Remove(conn); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				conn.Close()
			} else {
				log.Printf("msg: %s", string(msg))
			}
		}
	}
}
