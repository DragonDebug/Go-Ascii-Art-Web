package pages

import (
    "asciiartweb/asciiart"
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Output string
}

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("./HTML/Home.html")
    if err != nil {
        http.Error(w, "Error parsing template", http.StatusInternalServerError)
        log.Println("Error parsing template:", err)
        return
    }

    if r.Method == http.MethodPost {
        // Get the input text from the form
        inputText := r.FormValue("inputText")
		banner:= r.FormValue("banner")
        log.Println("Received input:", inputText, banner)

        // Check if input is coming through
        if inputText == "" {
            log.Println("No input text received")
        }

        // Convert input to ASCII Art (replace this with your actual ASCII art logic)
        output := asciiart.AsciiArt(inputText, banner)
        log.Println("Generated ASCII Art:", output)

        // Create a PageData instance and pass the output
        data := PageData{
            Output: output,
        }

        // Render the result using the template
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Error executing template", http.StatusInternalServerError)
            log.Println("Error executing template:", err)
            return
        }
    } else {
        // Render the page without any ASCII art output initially
        err = tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, "Error executing template", http.StatusInternalServerError)
            log.Println("Error executing template:", err)
            return
        }
    }
}
