package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "localhost:5000", "Server listening address")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	log.Println("Listening on", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}

	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}

	if ip == "" {
		ip = r.RemoteAddr
	}

	if ip == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	w.Write([]byte(ip))
}
