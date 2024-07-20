package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	art "ascii-art-web/ascii-art"
)

type Data struct {
	Filename string
	Input    string
	Result   string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", GetAsciiForm)
	http.HandleFunc("/ascii-art", PostAsciiArt)
	fmt.Println("SUCCESS!! listen to server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func GetAsciiForm(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Printf("Error executing template: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	default:
		http.NotFound(w, r)
	}
}

func PostAsciiArt(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request on /ascii-art route\n", r.Method)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("input")
	banner := r.FormValue("filename")

	if text == "" || banner == "" {
		http.Error(w, "400 Missing text or banner", http.StatusBadRequest)
		return
	}

	result, err := art.AsciiArt(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	resultData := &Data{
		Filename: banner,
		Input:    text,
		Result:   result,
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, resultData)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
