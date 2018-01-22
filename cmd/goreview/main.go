package main

import (
	"flag"
	"fmt"
	"net/http"

	hdlr "github.com/kelvintaywl/goreview/handler"
)

var serverPort int

func init() {
	flag.IntVar(&serverPort, "port", 9999, "port to expose for server")
}

func main() {
	http.HandleFunc("/", hdlr.IndexHandler)
	http.HandleFunc("/hooks", hdlr.HookHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}
