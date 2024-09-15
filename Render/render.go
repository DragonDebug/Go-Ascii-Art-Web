package render

import (
	"html/template"
	"log"
	"net/http"
)

// Render function to render the template
func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {
	parsedTemplate, err := template.ParseFiles("./html/" + tmpl)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
