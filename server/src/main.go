package main

import (
	"log"
	"net/http"
)

func main() {
	q := &stub{}
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/sura", suraHandler(q))
	log.Fatalln(server.ListenAndServe())
}
