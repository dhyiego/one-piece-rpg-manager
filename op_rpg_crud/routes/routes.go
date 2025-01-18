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

	// Devil Fruit routes
	router.HandleFunc("/devil-fruits", controllers.CreateDevilFruit).Methods("POST")
	router.HandleFunc("/devil-fruits", controllers.GetAllDevilFruits).Methods("GET")
	router.HandleFunc("/devil-fruits/{id}", controllers.GetDevilFruitByID).Methods("GET")

	return router
}
