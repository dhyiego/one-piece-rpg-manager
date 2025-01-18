package routes

import (
	"op_rpg_crud/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Character routes
	router.HandleFunc("/characters", controllers.CreateCharacter).Methods("POST")
	router.HandleFunc("/characters", controllers.GetAllCharacters).Methods("GET")
	router.HandleFunc("/characters/{id}", controllers.GetCharacterByID).Methods("GET")
	router.HandleFunc("/characters/{id}", controllers.UpdateCharacter).Methods("PUT")
	router.HandleFunc("/characters/{id}", controllers.DeleteCharacter).Methods("DELETE")

	// Attribute routes
	router.HandleFunc("/attributes", controllers.CreateAttribute).Methods("POST")
	router.HandleFunc("/attributes", controllers.GetAllAttributes).Methods("GET")
	router.HandleFunc("/attributes/{id}", controllers.GetAttributeByID).Methods("GET")

	// Fruit routes
	router.HandleFunc("/fruits", controllers.CreateFruit).Methods("POST")
	router.HandleFunc("/fruits", controllers.GetAllFruits).Methods("GET")
	router.HandleFunc("/fruits/{id}", controllers.GetFruitByID).Methods("GET")

	return router
}
