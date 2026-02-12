package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

}

func (s *Server) createPersonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreatePersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	now := time.Now()
	person := Person{
		ID:        1,
		Name:      req.Name,
		Document:  req.Document,
		CreatedAt: now,
		UpdatedAt: now,
	}
	book := Book{
		ID: 1,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		&CreatePersonResponse{
			ID:        person.ID,
			BookID:    book.ID,
			Name:      person.Name,
			Document:  person.Document,
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
		},
	)

}
