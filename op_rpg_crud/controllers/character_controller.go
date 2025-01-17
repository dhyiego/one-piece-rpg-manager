package controllers

import (
	"encoding/json"
	"net/http"
	"op_rpg_crud/config"
	"op_rpg_crud/models"

	"github.com/gorilla/mux"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var character models.Character
	json.NewDecoder(r.Body).Decode(&character)

	if result := config.DB.Create(&character); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(character)
}

func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	var characters []models.Character
	config.DB.Find(&characters)
	json.NewEncoder(w).Encode(characters)
}

func GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var character models.Character
	if result := config.DB.First(&character, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(character)
}

func UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var character models.Character
	if result := config.DB.First(&character, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&character)
	config.DB.Save(&character)
	json.NewEncoder(w).Encode(character)
}

func DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	config.DB.Delete(&models.Character{}, id)
	w.WriteHeader(http.StatusNoContent)
}
