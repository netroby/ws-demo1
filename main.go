package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	fmt.Println("Listening on :18017")
	err := http.ListenAndServe(":18017", nil)
	if err != nil {
		panic("ListenAndServe:" + err.Error())
	}
}
