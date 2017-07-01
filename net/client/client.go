package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var host string = "localhost:8080"
	var err error

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			fmt.Println("Usage:", os.Args[0], "[hostname:port]\nDefault:", host)
			return
		}
		host = os.Args[1]
	}
	fmt.Println("Get url:", "http://"+host)
	resp, err := http.Get("http://" + host)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
