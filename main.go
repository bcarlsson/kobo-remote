package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = "10000"

func left(w http.ResponseWriter, r *http.Request) {
	bin, err := base64.RawStdEncoding.DecodeString(leftTouchEvent)
	if err != nil {
		log.Fatal(err)
	}

	sendTouchEvent(bin)

}

func right(w http.ResponseWriter, r *http.Request) {
	bin, err := base64.RawStdEncoding.DecodeString(rightTouchEvent)
	if err != nil {
		log.Fatal(err)
	}

	sendTouchEvent(bin)
}

func sendTouchEvent(touchEvent []byte) {
	f, err := os.OpenFile("/dev/input/event1",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.Write(touchEvent); err != nil {
		log.Println(err)
	}
}

func handleRequests() {
	http.HandleFunc("/left", left)
	http.HandleFunc("/right", right)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	handleRequests()
}
