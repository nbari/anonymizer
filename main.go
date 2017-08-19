package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	director := func(req *http.Request) {
		req = r
		req.URL.Scheme = "https"
		req.URL.Host = r.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}

func start(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// for https
		if r.Method == "CONNECT" {
			// TODO
			fmt.Printf("r.Method = %+v\n", r.Method)
			return
		}
		if !r.URL.IsAbs() {
			fmt.Println("export http_proxy=localhost:8080")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func randomizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgents := []string{"foo", "bar", "todo"}
		// Select a random endpoint
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(userAgents))
		ua := userAgents[i]
		r.Header.Set("User-Agent", ua)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.Handle("*", middleware.New(start, randomizer).ThenFunc(Proxy))
	log.Fatal(http.ListenAndServe(":8080", router))
}
