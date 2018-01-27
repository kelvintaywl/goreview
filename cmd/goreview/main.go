package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	hdlr "github.com/kelvintaywl/goreview/handler"
)

var serverPort int

func init() {
	p := os.Getenv("PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		port = 9999
	}
	flag.IntVar(&serverPort, "port", port, "port to expose for server")
}

func main() {
	http.HandleFunc("/", hdlr.IndexHandler)
	http.HandleFunc("/hooks", hdlr.HookHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}
