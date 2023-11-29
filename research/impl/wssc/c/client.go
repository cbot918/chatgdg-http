package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	// Send WebSocket upgrade request
	fmt.Fprint(conn, "GET / HTTP/1.1\r\n")
	fmt.Fprint(conn, "Host: localhost:8889\r\n")
	fmt.Fprint(conn, "Upgrade: websocket\r\n")
	fmt.Fprint(conn, "Connection: Upgrade\r\n")
	key := "dGhlIHNhbXBsZSBub25jZQ=="
	fmt.Fprint(conn, "Sec-WebSocket-Key: "+key+"\r\n")
	fmt.Fprint(conn, "Sec-WebSocket-Version: 13\r\n")
	fmt.Fprint(conn, "\r\n")

	// Read response
	tp := textproto.NewReader(bufio.NewReader(conn))
	line, err := tp.ReadLine()
	if err != nil {
		fmt.Println("Error reading line:", err)
		return
	}

	if !strings.Contains(line, "101") {
		fmt.Println("Failed to upgrade:", line)
		return
	}

	// Read and print headers (optional)
	headers, err := tp.ReadMIMEHeader()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return
	}

	for k, v := range headers {
		fmt.Println(k+":", v)
	}

	// Verify Sec-WebSocket-Accept
	h := sha1.New()
	h.Write([]byte(key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	expectedAccept := base64.StdEncoding.EncodeToString(h.Sum(nil))
	if headers.Get("Sec-WebSocket-Accept") != expectedAccept {
		fmt.Println("Invalid Sec-WebSocket-Accept")
		return
	}

	fmt.Println("WebSocket connection established")

	// Now you can read/write from/to conn for WebSocket communication
	// Remember, WebSocket data framing needs to be handled for communication
}
