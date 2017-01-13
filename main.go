package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 172.17.0.2 - - [12/Jan/2017 20:21:40] "GET /README.md HTTP/1.1" 200 -
		log.Printf("%s %s", getIP(r), r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Serving HTTP on 0.0.0.0 port 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", LogHandler(http.FileServer(http.Dir(".")))))
}

func getIP(req *http.Request) string {

	// try proxy friendly headers
	for _, header := range []string{"X-Real-Ip", "X-Forwarded-For"} {
		realIP, ok := req.Header[header]
		if ok {
			return cleanIP(realIP[0])
		}
	}

	// fall back to the remote address
	return cleanIP(req.RemoteAddr)
}

func cleanIP(address string) string {
	address = strings.TrimPrefix(address, "::ffff:")
	return strings.Split(address, ":")[0]
}
