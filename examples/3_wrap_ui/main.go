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

	f := NewFrame()

	_, rw, err := Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

	// handle inMessage
	buf := make([]byte, 4096)
	_, err = rw.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	inMessage := f.DecodeFrame(buf)
	fmt.Println(string(inMessage))

	// handle outMessage
	outMessage := f.EncodeFrame(inMessage)
	_, err = rw.Write(outMessage)
	if err != nil {
		log.Fatal(err)
	}
	err = rw.Flush()
	if err != nil {
		log.Fatal(err)
	}

}
