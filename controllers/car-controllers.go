package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gustasousaghs/carros-pcd-service/models"
	"github.com/gustasousaghs/carros-pcd-service/service"
)

func CreateCarController(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "error ao codifocar json", http.StatusBadRequest)
		return
	}
	if err := service.CreateCar(&car); err != nil {
		http.Error(w, "Erro ao criar carro", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(car)
}
func UpdateCarController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	if err := service.UpdateCar(id, car); err != nil {
		http.Error(w, "Erro ao atualizar carro", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Carro atualizado com sucesso"))
}

func DeleteCarController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := service.DeleteCar(id); err != nil {
		http.Error(w, "Erro ao deletar carro", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Carro deletado com sucesso"))
}
func GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := service.GetAllCars()
	if err != nil {
		http.Error(w, "erro ao chamar carros", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&cars)
}
func GetCarById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	cars, err := service.GetCarById(id)
	if err != nil {
		http.Error(w, "erro ao pegar carro", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
