package pages

import (
	render "asciiartweb/Render"
	art "asciiartweb/asciiart"
	"log"
	"net/http"
)

type PageData struct {
	Output string
}

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		render.RenderTemplate(w, "home.html", nil)

	case http.MethodPost:
		// Get the input text from the form
		inputText := r.FormValue("inputText")
		banner := r.FormValue("banner")
		log.Println("Received input:", inputText, banner)

		// Convert input to ASCII Art
		output := art.AsciiArt(inputText, banner)
		log.Print("Generated ASCII Art:\n",output)

		// Create a PageData instance and pass the output
		data := PageData{
			Output: output,
		}

		// Render the result using the template
		render.RenderTemplate(w, "home.html", data)
	}
}
