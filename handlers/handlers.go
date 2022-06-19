package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Lightbulb struct {
	Name string
	On   bool
}

type Repository interface {
	GetLightbulbs(ctx context.Context) (map[string]bool, error)
	GetLightbulbById(ctx context.Context, name string) (bool, error)
	CreateLightbulbs(ctx context.Context, name string, value bool) error
	UpdateLightbulb(ctx context.Context, name string, value bool) error
	DeleteLightbulb(ctx context.Context, name string) error
}

func HandleLightbulbs(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetLightbulbs(repo)
			return
		case http.MethodPost:
			CreateLightbulbs(repo)
			return
		case http.MethodPut:
			SwitchLightbulbsState(repo)
			return
		case http.MethodDelete:
			DeleteLightbulb(repo)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func GetLightbulbs(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		lightbulbs, err := repo.GetLightbulbs(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}
		json.NewEncoder(w).Encode(lightbulbs)
	}
}

func CreateLightbulbs(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lb Lightbulb
		err := json.NewDecoder(r.Body).Decode(&lb)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err.Error())
			return
		}
		err = repo.CreateLightbulbs(r.Context(), lb.Name, lb.On)
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}

		lightbulbs, err := repo.GetLightbulbs(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lightbulbs)
	}
}

func SwitchLightbulbsState(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		currentStatus, err := repo.GetLightbulbById(r.Context(), name)
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}

		err = repo.UpdateLightbulb(r.Context(), name, !currentStatus)
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}

		lightbulbs, err := repo.GetLightbulbs(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}

		json.NewEncoder(w).Encode(lightbulbs)
	}
}

func DeleteLightbulb(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := repo.DeleteLightbulb(r.Context(), name)
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}

		lightbulbs, err := repo.GetLightbulbs(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			fmt.Println(err.Error())
			return
		}
		json.NewEncoder(w).Encode(lightbulbs)
	}
}
