package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 172.17.0.2 - - [12/Jan/2017 20:21:40] "GET /README.md HTTP/1.1" 200 -
		log.Printf("- %s - %s", getIP(r), r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := flag.Int("port", 8000, "port to listen on")
	flag.Parse()

	fmt.Printf("Serving HTTP on 0.0.0.0 port %d ...\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), LogHandler(http.FileServer(http.Dir(".")))))
}

func getIP(req *http.Request) string {

	// try proxy friendly headers
	for _, header := range []string{"X-Real-Ip", "X-Forwarded-For"} {
		realIP, ok := req.Header[header]
		if ok {
			return realIP[0]
		}
	}

	// fall back to the remote address
	return req.RemoteAddr
}
