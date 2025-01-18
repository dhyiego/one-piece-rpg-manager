package controllers

import (
	"encoding/json"
	"net/http"
	"op_rpg_crud/config"
	"op_rpg_crud/models"
	"op_rpg_crud/op_rpg_crud/models"

	"github.com/gorilla/mux"
)

func CreateFruit(w http.ResponseWriter, r *http.Request) {
	var fruit models.Fruit
	json.NewDecoder(r.Body).Decode(&fruit)

	if result := config.DB.Create(&fruit); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fruit)
}

func GetAllFruits(w http.ResponseWriter, r *http.Request) {
	var fruits []models.Fruit
	config.DB.Find(&fruits)
	json.NewEncoder(w).Encode(fruits)
}

func GetFruitByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var fruit models.Fruit
	if result := config.DB.First(&fruit, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(fruit)
}
