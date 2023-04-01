package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/atotto/clipboard"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":9999", "listen address")
}

func main() {
	if err := mainE(); err != nil {
		log.Fatal(err)
	}
}

func mainE() error {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		go func(s string) {
			if err := clipboard.WriteAll(s); err != nil {
				log.Println(err)
				return
			}
			log.Printf("copied %d bytes, content=%#v", len(s), s)
		}(string(b))
    w.WriteHeader(http.StatusAccepted)
	})

	log.Printf("listening on %s", addr)
	return http.ListenAndServe(addr, mux)
}
