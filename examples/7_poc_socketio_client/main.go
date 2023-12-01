package main

import (
	"chatgdg-http/examples/6_add-gin/ws"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/ws", handleWS)

	err := router.Run(port)
	if err != nil {
		log.Fatal(err)
	}

	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}

const magicWebSocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func handleWS(c *gin.Context) {
	fmt.Println("in handleWS")
	w := c.Writer
	r := c.Request
	chat.UserCount += 1

	clientID := r.URL.Query().Get("name")

	fmt.Println("42")

	conn, rw, err := ws.Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conn)

	fmt.Println(rw)

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
