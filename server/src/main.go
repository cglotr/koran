package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/handlers"
	"github.com/cglotr/koran/server/src/middlewares"
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
	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/sura", handlers.SuraHandler(mysql)).Methods(http.MethodGet)
	r.HandleFunc("/translation", handlers.TranslationHandler(mysql)).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}/read", handlers.UserReadHandler()).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.RootHandler()).Methods(http.MethodGet)

	r.HandleFunc("/auth", handlers.AuthHandler(client, mysql)).Methods(http.MethodPost)
	r.HandleFunc("/auth/{id}/invalidate", handlers.AuthInvalidateHandler(client, mysql)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}/read", handlers.UserReadPostHandler()).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}/unread", handlers.UserUnreadPostHandler()).Methods(http.MethodPost)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}
