package handler

import (
	"encoding/json"
	"net/http"

	"repository" // want "handler must not include repository"
	"service"
)

// DeleteHandler is Dummay for test
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	repo := repository.UserRepository()
	u := repo.Get(req.ID)
	service.DeleteUser(u)
}
