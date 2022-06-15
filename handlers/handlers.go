package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luisadiniz/homecontroller/repositories"
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

	lightbulbs, err := repositories.GetLightbulbs(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}
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
	err = repositories.CreateLightbulbs(r.Context(), lb.Name, lb.On)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	lightbulbs, err := repositories.GetLightbulbs(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lightbulbs)
}

func SwitchLightbulbsState(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	currentStatus, err := repositories.GetLightbulbById(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	err = repositories.UpdateLightbulb(r.Context(), name, !currentStatus)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	lightbulbs, err := repositories.GetLightbulbs(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	json.NewEncoder(w).Encode(lightbulbs)
}

func DeleteLightbulb(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := repositories.DeleteLightbulb(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	lightbulbs, err := repositories.GetLightbulbs(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(lightbulbs)
}
