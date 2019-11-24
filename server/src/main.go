package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	q := &stub{}
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/sura", getHandler(suraHandler(q)))
	http.HandleFunc("/", getHandler(rootHandler()))
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}
