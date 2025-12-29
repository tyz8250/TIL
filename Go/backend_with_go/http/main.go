package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s, remote=%s", r.Method, r.RemoteAddr)
		fmt.Fprintln(w, "hello")
	})

	http.ListenAndServe(":8080", nil)

}

// go run main.go
// curl -v http://localhost:8080/hello
//answer

// * Host localhost:8080 was resolved.
// * IPv6: ::1
// * IPv4: 127.0.0.1
// *   Trying [::1]:8080...
// * Connected to localhost (::1) port 8080
// > GET /hello HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/8.7.1
// > Accept: */*
// >
// * Request completely sent off
// < HTTP/1.1 200 OK
// < Date: Mon, 29 Dec 2025 12:38:19 GMT
// < Content-Length: 6
// < Content-Type: text/plain; charset=utf-8
// <
// hello
// * Connection #0 to host localhost left intact
// `
