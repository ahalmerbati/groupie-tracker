package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
)

func renderErrorTemplate(w http.ResponseWriter, code int, data models.ErrorData) {
	templatePath := fmt.Sprintf("templates/errors/%d.html", code)

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error: Cannot parse required error template %s: %v", templatePath, err)
		
		// If the Error template is not parsing this is the fallback (calling the http.error method)
		http.Error(w, fmt.Sprintf("%d %s", code, data.StatusText), code)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing error template %s: %v", templatePath, err)
		http.Error(w, fmt.Sprintf("%d %s", code, data.StatusText), code)
		return
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		log.Printf("Error parsing templates %s: %v", tmpl, err)
		ErrorHandler(w, nil, http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
