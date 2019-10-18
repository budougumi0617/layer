package nested

import (
	"encoding/json"
	"net/http"

	rn "repository/nested" // want "handler/nested must not include repository/nested"
	sn "service/nested"
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
	repo := rn.UserRepository()
	u := repo.Get(req.ID)
	sn.DeleteUser(u)
}
