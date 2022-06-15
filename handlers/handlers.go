package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	lightbulbs = map[string]bool{}
)

type Lightbulb struct {
	Name string
	On   bool
}

func HandleLightbulbs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetLightbulbs(w, r)
		return
	case http.MethodPost:
		CreateLightbulbs(w, r)
		return
	case http.MethodPut:
		SwitchLightbulbsState(w, r)
		return
	case http.MethodDelete:
		DeleteLightbulb(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func GetLightbulbs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lightbulbs)
}

func CreateLightbulbs(w http.ResponseWriter, r *http.Request) {
	var lb Lightbulb
	err := json.NewDecoder(r.Body).Decode(&lb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	lightbulbs[lb.Name] = lb.On
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lightbulbs)
}

func SwitchLightbulbsState(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	lightbulbs[name] = !lightbulbs[name]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lightbulbs)
}

func DeleteLightbulb(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	delete(lightbulbs, name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lightbulbs)
}
