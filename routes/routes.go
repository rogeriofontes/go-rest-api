package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rogeriofontes/go-rest-api/controllers"
	"github.com/rogeriofontes/go-rest-api/middleware"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	router.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	router.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	router.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	router.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	//log.Fatal(http.ListenAndServe(":8000", r))
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
