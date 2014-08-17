package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	bind = flag.String("bind", ":8080", "HTTP listener address")
)

func init() {
	flag.Parse()
}

func handler(rw http.ResponseWriter, req *http.Request) {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	xff := req.Header.Get("X-Forwarded-For")
	rw.Header().Set("Content-Type", "text/plain")
	if len(xff) > 0 {
		fmt.Fprintf(rw, "%s, %s\n", xff, host)
	} else {
		fmt.Fprintf(rw, "%s\n", host)
	}
}

func main() {
	log.Fatalln(http.ListenAndServe(*bind, http.HandlerFunc(handler)))
}
