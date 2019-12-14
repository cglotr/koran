package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql, err := database.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = mysql.Db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/sura", handlers.GetHandler(handlers.SuraHandler(mysql)))
	http.HandleFunc("/", handlers.GetHandler(handlers.RootHandler()))
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}
