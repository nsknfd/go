package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var num int

func main() {
	num = 0
	var port int = 8080
	var err error

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			fmt.Println("Usage:", os.Args[0], "[port]\nDefault port: 8080")
			return
		}
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid port:", os.Args[1])
			fmt.Println("Use default port(8080)")
			port = 8080
		}
	}
	fmt.Println("Listen port:", port)

	http.HandleFunc("/", Hello)
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[", num, "] Request From: ", req.RemoteAddr)
	num = num + 1
	w.Write([]byte("Hello!\n"))
	w.Write([]byte(req.RemoteAddr))
}
