package main

import (
	"log"
	"net/http"
	"golang.org/x/net/webdav"
	"fmt"
	"flag"
	"net"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Set server port.")
	flag.Parse()
}

func main() {
	srv := &webdav.Handler{
		FileSystem: webdav.NewMemFS(),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			log.Printf("%s Incoming %s method from: %s ", r.URL, r.Method, ip)
		},
	}

	hostPort := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Initiating test server listening at %s", hostPort)

	http.Handle("/", srv)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
