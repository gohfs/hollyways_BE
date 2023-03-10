package main

import (
	"fmt"
	"hollyways/database"
	"hollyways/pkg/mysql"
	"hollyways/routes"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// DB CONNECT
	mysql.DatabaseInit()

	// MIGRATION
	database.RunMigration()

	r := mux.NewRouter()

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-Width", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())
	
	var PORT = os.Getenv('PORT')

	fmt.Println("SERVER IS ONLINE 🚀")
	http.ListenAndServe(":" +PORT, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
