package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
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
	http.HandleFunc("/auth", handlers.GetHandler(handlers.AuthHandler(client, mysql)))
	http.HandleFunc("/sura", handlers.GetHandler(handlers.SuraHandler(mysql)))
	http.HandleFunc("/translation", handlers.GetHandler(handlers.TranslationHandler(mysql)))
	http.HandleFunc("/", handlers.GetHandler(handlers.RootHandler()))
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}
