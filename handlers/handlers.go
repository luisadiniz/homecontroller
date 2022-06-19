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
	Get(ctx context.Context) (map[string]bool, error)
	GetById(ctx context.Context, name string) (bool, error)
	Create(ctx context.Context, name string, value bool) error
	Update(ctx context.Context, name string, value bool) error
	Delete(ctx context.Context, name string) error
}

func HandleLightbulbs(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Get(w, r, repo)
			return
		case http.MethodPost:
			Create(w, r, repo)
			return
		case http.MethodPut:
			SwitchState(w, r, repo)
			return
		case http.MethodDelete:
			Delete(w, r, repo)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func Get(w http.ResponseWriter, r *http.Request, repo Repository) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	lightbulbs, err := repo.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(lightbulbs)
}

func Create(w http.ResponseWriter, r *http.Request, repo Repository) {
	var lb Lightbulb
	err := json.NewDecoder(r.Body).Decode(&lb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	err = repo.Create(r.Context(), lb.Name, lb.On)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	lightbulbs, err := repo.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lightbulbs)
}

func SwitchState(w http.ResponseWriter, r *http.Request, repo Repository) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	currentStatus, err := repo.GetById(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	err = repo.Update(r.Context(), name, !currentStatus)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	lightbulbs, err := repo.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	json.NewEncoder(w).Encode(lightbulbs)
}

func Delete(w http.ResponseWriter, r *http.Request, repo Repository) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := repo.Delete(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}

	lightbulbs, err := repo.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(lightbulbs)
}
