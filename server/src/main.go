package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/handlers"
	"github.com/gorilla/mux"

	firebase "firebase.google.com/go"
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
	r := mux.NewRouter()

	r.HandleFunc("/sura", handlers.GetHandler(handlers.SuraHandler(mysql)))
	r.HandleFunc("/translation", handlers.GetHandler(handlers.TranslationHandler(mysql)))
	r.HandleFunc("/", handlers.GetHandler(handlers.RootHandler()))

	r.HandleFunc("/auth", handlers.PostHandler(handlers.AuthHandler(client, mysql)))
	r.HandleFunc("/auth/{id}/invalidate", handlers.PostHandler(handlers.AuthInvalidateHandler(client, mysql)))

	server := http.Server{
		Addr: ":8080",
	}
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}
