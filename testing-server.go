package main

import (
	"net/http"
	"strconv"

	"github.com/integrii/flaggy"
)

func main() {
	flaggy.SetName("testing-server")
	flaggy.SetDescription("HTTP testing server which responds to requests as configured.")

	var path = "/"
	flaggy.String(&path, "x", "path", "HTTP path pattern on which to listen. Pattern follows the same rules as the one documented in golang's http.ServceMux.")

	var port = 8080
	flaggy.Int(&port, "p", "port", "Port on which to listen.")

	var responseCode = 200
	flaggy.Int(&responseCode, "r", "response-code", "HTTP response code for all requests.")

	var certFile = ""
	flaggy.String(&certFile, "c", "tls-cert-file", "Certificate for TLS. Requires -k/--key to also be provided.")

	var keyFile = ""
	flaggy.String(&keyFile, "k", "tls-key-file", "Key for TLS. Requires -c/--cert to also be provided.")

	flaggy.Parse()

	if certFile != "" && keyFile == "" {
		println("TLS certificate was specified, but no key was.")
	} else if keyFile != "" && certFile == "" {
		println("TLS key was specified, but no certificate was.")
	}

	http.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(responseCode)
	})

	var addr = ":" + strconv.Itoa(port)

	if certFile != "" && keyFile != "" {
		http.ListenAndServeTLS(addr, certFile, keyFile, nil)
	} else {
		http.ListenAndServe(addr, nil)
	}
}
