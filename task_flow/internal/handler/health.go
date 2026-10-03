package handler

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, http.StatusOK, true, "Health is OK", nil)
}
