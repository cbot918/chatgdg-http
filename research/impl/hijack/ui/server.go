package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	port               = ":8889"
	magicWebSocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

func main() {
	http.HandleFunc("/", Upgrade)
	fmt.Println(port)
	http.ListenAndServe(port, nil)
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	upgrade(w, r)
}

func upgrade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in")
	key := r.Header.Get("Sec-WebSocket-Key")
	if key == "" {
		http.Error(w, "Missing Sec-WebSocket-Key header", http.StatusBadRequest)
		return
	}

	fmt.Println("key:", key)

	h := sha1.New()
	h.Write([]byte(key + magicWebSocketGUID))
	accept := base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Println("accept:", accept)

	w.Header().Set("Upgrade", "websocket")
	w.Header().Set("Connection", "Upgrade")
	w.Header().Set("Sec-WebSocket-Accept", accept)

	// Log the response headers
	for name, values := range w.Header() {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	w.WriteHeader(http.StatusSwitchingProtocols)

}
