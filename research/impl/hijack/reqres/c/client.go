package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8889")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#+v", resp)
	// printJSON(resp.Header)
}

func printJSON(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("marshal indent failed")
		return
	}
	fmt.Println(string(json))
}
