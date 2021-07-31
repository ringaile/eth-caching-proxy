package handler

import (
	"net/http"
)

func GetHealthz(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, nil)
}
