package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (net.Conn, *bufio.ReadWriter, error) {
	return upgrade(w, r)
}

func upgrade(w http.ResponseWriter, r *http.Request) (net.Conn, *bufio.ReadWriter, error) {

	// get websocket key from header
	key := r.Header.Get("Sec-WebSocket-Key")
	if key == "" {
		http.Error(w, "Missing Sec-WebSocket-Key header", http.StatusBadRequest)
		return nil, nil, fmt.Errorf("missing sec-webwocket-key header")
	}

	// prepare response key
	h := sha1.New()
	h.Write([]byte(key + magicWebSocketGUID))
	accept := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// add header
	w.Header().Set("Upgrade", "websocket")
	w.Header().Set("Connection", "Upgrade")
	w.Header().Set("Sec-WebSocket-Accept", accept)
	w.WriteHeader(http.StatusSwitchingProtocols)

	// hijack connection from http request
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Web server does not support hijacking", http.StatusInternalServerError)
		return nil, nil, fmt.Errorf("web server does not support hijacking")
	}

	conn, rw, err := hj.Hijack()
	if err != nil {
		http.Error(w, "Hijacking failed: "+err.Error(), http.StatusInternalServerError)
		return nil, nil, fmt.Errorf("hijacking failed")
	}

	return conn, rw, nil
}
