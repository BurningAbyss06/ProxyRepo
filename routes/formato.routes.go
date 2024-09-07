package routes

import (
	"encoding/json"
	"net/http"

	"github.com/BurningAbyss06/ProxyRepo/db"
	"github.com/BurningAbyss06/ProxyRepo/models"
	"github.com/gorilla/mux"
)

func GetFormatosHandler(w http.ResponseWriter, r *http.Request) {
	var formato []models.Formato
	temp := db.DB.Find(&formato)
	err := temp.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&formato)
}

func GetFormatoHandler(w http.ResponseWriter, r *http.Request) {
	var formato models.Formato
	params := mux.Vars(r)
	db.DB.First(&formato, params["id"])

	if formato.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Formato no Encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&formato)
}

func PostFormatoHandler(w http.ResponseWriter, r *http.Request) {
	var formato models.Formato
	json.NewDecoder(r.Body).Decode(&formato)

	temp := db.DB.Create(&formato)
	err := temp.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&formato)
}

func DeleteFormatoHandler(w http.ResponseWriter, r *http.Request) {
	var formato models.Formato
	params := mux.Vars(r)

	db.DB.First(&formato, params["id"])
	if formato.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Formato no Encontrado"))
		return
	}
	db.DB.Delete(&formato)
	w.WriteHeader(http.StatusOK)

}

func TrueDeleteFormatoHandler(w http.ResponseWriter, r *http.Request) {
	var formato models.Formato
	params := mux.Vars(r)

	db.DB.First(&formato, params["id"])
	if formato.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Formato no Encontrado"))
		return
	}
	db.DB.Unscoped().Delete(&formato)
	w.WriteHeader(http.StatusOK)

}
