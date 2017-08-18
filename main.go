package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
)

//func Proxy(w http.ResponseWriter, r *http.Request) *httputil.ReverseProxy {
func Proxy() *httputil.ReverseProxy {
	// TODO
	proxy := &httputil.ReverseProxy{}
	return proxy
}

func randomizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "CONNECT" {
			// TODO
			fmt.Printf("r.Method = %+v\n", r.Method)
			return
		}
		if r.Method == "GET" {
			// TODO
			fmt.Printf("r.Method = %+v\n", r.Method)
			return
		}
		if !r.URL.IsAbs() {
			fmt.Printf("r.URL = %+v\n", r.URL)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.Handle("*", middleware.New(randomizer).Then(Proxy()))
	log.Fatal(http.ListenAndServe(":8080", router))
}
