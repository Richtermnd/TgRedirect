package main

import (
	"log"
	"net/http"
	"strings"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	cutted, _ := strings.CutPrefix(r.URL.String(), "/tg")
	http.Redirect(w, r, "tg:/"+cutted, http.StatusFound)
}

func main() {
	http.HandleFunc("/tg/proxy", Redirect)
	log.Println("Start server")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
