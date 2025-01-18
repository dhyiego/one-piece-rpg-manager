package controllers

import (
	"encoding/json"
	"net/http"
	"op_rpg_crud/config"
	"op_rpg_crud/models"

	"github.com/gorilla/mux"
)

func CreateAttribute(w http.ResponseWriter, r *http.Request) {
	var attribute models.Attribute
	json.NewDecoder(r.Body).Decode(&attribute)

	// Calculate Charisma and Defense
	attribute.Charisma = (attribute.Intelligence + attribute.Willpower) / 2
	attribute.Defense = attribute.Vitality * 2

	if result := config.DB.Create(&attribute); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attribute)
}

func GetAllAttributes(w http.ResponseWriter, r *http.Request) {
	var attributes []models.Attribute
	config.DB.Find(&attributes)
	json.NewEncoder(w).Encode(attributes)
}

func GetAttributeByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var attribute models.Attribute
	if result := config.DB.First(&attribute, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(attribute)
}
