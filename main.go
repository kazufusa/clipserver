package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
)

var _ http.Handler = (*handler)(nil)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	clipboard.WriteAll(string(body))
	log.Println(string(body))
}

var port = "8080"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help":
			fmt.Println("clipserver serves posted data to clipboard")
			fmt.Println("")
			fmt.Println("optional arguments")
			fmt.Println("  -h, --help            show this help message and exit")
			fmt.Println("  -p PORT, --port PORT  specify port [default:8080]")
			return
		case "-p", "--port":
			port = os.Args[2]
		}
	}
	http.Handle("/", handler{})
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}
