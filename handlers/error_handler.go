package handlers

import (
	"groupie-tracker/models"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int) {
	statusText := http.StatusText(code)

	data := models.ErrorData{
		StatusCode: code,
		StatusText: statusText,
	}

	renderErrorTemplate(w, code, data)
}
