package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
)

var (
	bind = flag.String("bind", ":80", "HTTP listener address")
	xff  = flag.Bool("xff", false, "expose X-Forwarded-For")
)

func init() {
	flag.Parse()
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	if !*xff {
		fmt.Fprintf(rw, "%s\n", host)
		return
	}
	v := req.Header.Get("X-Forwarded-For")
	if len(v) > 0 {
		v += ", "
	}
	fmt.Fprintf(rw, "%s%s\n", v, host)
}

func main() {
	log.Fatalln(http.ListenAndServe(*bind, http.HandlerFunc(handler)))
}
