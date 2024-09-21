package main

import (
	"net/http"

	"github.com/BurningAbyss06/ProxyRepo/db"
	"github.com/BurningAbyss06/ProxyRepo/models"
	"github.com/BurningAbyss06/ProxyRepo/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Formato{})
	db.DB.AutoMigrate(models.F_D{})
	db.DB.AutoMigrate(models.Deck{})
	db.DB.AutoMigrate(models.Decklist{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/formato/Get", routes.GetFormatosHandler).Methods("GET")
	router.HandleFunc("/formato/Get/{id}", routes.GetFormatoHandler).Methods("GET")
	router.HandleFunc("/formato/Post", routes.PostFormatoHandler).Methods("POST")
	router.HandleFunc("/formato/Delete/{id}", routes.DeleteFormatoHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
