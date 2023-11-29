package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const port = ":8787"

func main() {

	http.HandleFunc("/ws", wsHandler)

	fmt.Println(port)
	http.ListenAndServe(port, nil)

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}

	// buf := make([]byte, 1024)
	// n, err := conn.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(buf[:n]))

	msg, _, err := wsutil.ReadClientData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(msg))
}
