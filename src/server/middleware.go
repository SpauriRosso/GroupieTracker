package Serveur

import (
	"groupie-tracker/src/shared"
	"net/http"
)

func ResultAllowed(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "POST" {
		shared.Prompt = r.FormValue("search")
		//s.formSubmitted = true
		//http.Redirect(w, r, "/result", http.StatusSeeOther)
		return true
	}
	return false
}
