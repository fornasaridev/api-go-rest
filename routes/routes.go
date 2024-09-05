package routes

import (
	"api-go-rest/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc(
		"/api/personalidades",
		controllers.TodasPersonalidades,
	)
	log.Fatal(http.ListenAndServe(":8080", r))

}
