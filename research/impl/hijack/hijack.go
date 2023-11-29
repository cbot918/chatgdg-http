package main

/*
GPT
can you write a simple natvie golang  websocket server without 3rd library  it can upgrade and decode the first message from client ?

*/
import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

const (
	port               = ":8888"
	magicWebSocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

var (
	err            error
	ErrBadHijacker = fmt.Errorf("given http.ResponseWriter is not a http.Hijacker")
)

func main() {
	http.HandleFunc("/ws", upgrade)
	fmt.Println(port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) {

	fmt.Println("in")

	key := r.Header.Get("Sec-WebSocket-Key")
	if key == "" {
		http.Error(w, "Missing Sec-WebSocket-Key header", http.StatusBadRequest)
		return
	}

	fmt.Println("key: ", key)

	h := sha1.New()
	h.Write([]byte(key + magicWebSocketGUID))
	accept := base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Println("accept: ", accept)

	w.Header().Set("Upgrade", "websocket")
	w.Header().Set("Connection", "Upgrade")
	w.Header().Set("Sec-WebSocket-Accept", accept)
	w.WriteHeader(http.StatusSwitchingProtocols)

	fmt.Fprintln(w)

	hj, ok := w.(http.Hijacker)
	if !ok {
		err = ErrBadHijacker
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	conn, buf, err := hj.Hijack()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Read the first WebSocket frame
	frame, err := bufio.NewReader(buf).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading frame:", err)
		return
	}

	// Process the frame to decode the message
	f := NewFrame()
	message := f.DecodeFrame(frame)
	// message, err := decodeFrame()
	// if err != nil {
	// 	fmt.Println("Error decoding frame:", err)
	// 	return
	// }

	fmt.Println("Received message:", string(message))

}
