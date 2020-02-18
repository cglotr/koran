package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/handlers"
	"github.com/cglotr/koran/server/src/middlewares"
	"github.com/cglotr/koran/server/src/user"
	"github.com/cglotr/koran/server/src/usermysql"
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

	userCRUD := getUserCRUD()

	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/sura", handlers.SuraHandler(mysql)).Methods(http.MethodGet)
	r.HandleFunc("/translation", handlers.TranslationHandler(mysql)).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}/read", handlers.UserReadHandler(mysql)).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.RootHandler()).Methods(http.MethodGet)

	r.HandleFunc("/auth", handlers.AuthHandler(client, userCRUD)).Methods(http.MethodPost)
	r.HandleFunc("/auth/{id}/invalidate", handlers.AuthInvalidateHandler(client, userCRUD)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}/read", handlers.UserReadPostHandler(mysql)).Methods(http.MethodPost)

	r.HandleFunc("/user/{id}/read", handlers.UserReadDeleteHandler(mysql)).Methods(http.MethodDelete)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if flag.Lookup("test.v") == nil {
		log.Fatalln(server.ListenAndServe())
	}
}

func getUserCRUD() user.CRUD {
	userMySQL, err := usermysql.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = userMySQL.Db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return userMySQL
}
