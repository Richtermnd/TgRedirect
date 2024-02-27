package main

import (
	"log"
	"net/http"
	"strings"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	cutted, _ := strings.CutPrefix(r.URL.String(), "/tg")
	new := "tg:/" + cutted
	log.Printf("Redirect from %s to %s", r.URL.String(), new)
	http.Redirect(w, r, new, http.StatusFound)
}

func main() {
	http.HandleFunc("/tg/proxy", Redirect)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	log.Println("Start server")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
