package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port               = ":8889"
	magicWebSocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

func main() {
	router := gin.Default()
	router.GET("/ws", handleWS)

	fmt.Println(port)
	http.ListenAndServe(port, nil)
}

func handleWS(c *gin.Context) {
	w := c.Writer
	r := c.Request

	_, _, err := Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

}
