// handlers.go

package handlers

import (
	"ascii-art-web/ascii-art"
	"html/template"
	"net/http"
	"os"
)

// HandleAsciiArt processes the form data for generating ASCII art.
func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST; return 405 Method Not Allowed otherwise
	if r.Method != http.MethodPost {
		renderErrorTemplate(w, http.StatusMethodNotAllowed, "errors/405.html")
		return
	}

	// Parse form data from the request
	err := r.ParseForm()
	if err != nil {
		// Handle form parsing errors by returning a 400 Bad Request
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Retrieve form values for art style and user text
	artStyle := r.FormValue("artstyle")
	userText := r.FormValue("text")

	// Validate that both art style and user text are provided
	if artStyle == "" || userText == "" {
		// Return 400 Bad Request if either value is missing
		http.Error(w, "Missing art style or text", http.StatusBadRequest)
		return
	}

	// Construct the path to the art style file
	artStylePath := "ascii-art/artstyles/" + artStyle + ".txt"

	// Check if the art style file exists
	if _, err := os.Stat(artStylePath); os.IsNotExist(err) {
		// Return 400 Bad Request if the art style file is invalid
		http.Error(w, "Invalid art style", http.StatusBadRequest)
		return
	}

	// Generate the ASCII art based on user text and selected art style
	asciiArtResult := ascii_art.AsciiArt(userText, artStylePath)

	// Prepare the data for rendering the template
	data := struct {
		ASCIIArtResult string
	}{
		ASCIIArtResult: asciiArtResult,
	}

	// Render the main template with the generated ASCII art
	tmplName := "index.html"
	renderTemplateWithData(w, tmplName, data)
}

// ServeTemplate serves the HTML template for the home page
func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
	}{
		Message: "Welcome to ASCII Art Web!",
	}
	renderTemplateWithData(w, "index.html", data)
}

// renderTemplateWithData renders the HTML template with the provided data
func renderTemplateWithData(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// renderErrorTemplate renders an error page
func renderErrorTemplate(w http.ResponseWriter, statusCode int, templatePath string) {
	http.Error(w, http.StatusText(statusCode), statusCode)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
