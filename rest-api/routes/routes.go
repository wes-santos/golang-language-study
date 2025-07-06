package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wes-santos/rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnAPersonality).Methods("Get")

	r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", r))
}
