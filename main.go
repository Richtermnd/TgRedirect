package main

// Server redirect from http://example.com/tg/proxy to tg://proxy

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	Port string
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	cutted, _ := strings.CutPrefix(r.URL.String(), "/tg/")
	new := "tg://" + cutted
	log.Printf("Redirect from %s to %s", r.URL.String(), new)
	http.Redirect(w, r, new, http.StatusFound)
}

func main() {
	// Load port from enviroment
	loadEnv()

	// Set handler
	http.HandleFunc("/tg/proxy", Redirect)

	// Start server
	log.Println("Start server")
	if err := http.ListenAndServe(":"+Port, nil); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	Port = os.Getenv("PORT")
	if Port == "" {
		log.Fatal("PORT env var not found")
	}
	if _, err := strconv.ParseInt(Port, 10, 64); err != nil {
		log.Fatalf("PORT %s is not integer", Port)
	}
}
