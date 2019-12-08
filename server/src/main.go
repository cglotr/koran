package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mysql, err := open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = mysql.db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/sura", getHandler(suraHandler(mysql)))
	http.HandleFunc("/", getHandler(rootHandler()))
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}
