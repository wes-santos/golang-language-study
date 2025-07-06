package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wes-santos/rest-api/database"
	"github.com/wes-santos/rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func ReturnAPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
