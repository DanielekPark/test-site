package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"html/template"
	"io"
	"log"
	"net/http"

	"github/strapi-admin/pkg/models"

	"github.com/joho/godotenv"
)

func checkPath(expectedPath string, w http.ResponseWriter, r *http.Request) error {

	if r.URL.Path != expectedPath {
		tmpl, err := template.ParseFiles("static/templates/404.html")
		if err != nil {
			return fmt.Errorf("failed to parse 404 template: %w", err)
		}
		if err := tmpl.Execute(w, nil); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	}
	return nil
}

func FetchEvents(url string, token string) models.EventsPageData {
	godotenv.Load()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//Parses JSON
	var result models.EventsPageData
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Unable to marshal JSON")
	}
	return result
}

func Index(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/", w, r); err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/about", w, r); err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("static/templates/about.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("BEARER_TOKEN")
	url := os.Getenv("EVENTS_ENDPOINT")
	ParsedData := FetchEvents(url, token).Data

	if err := checkPath("/events", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/events.html")

	if err := tmpl.Execute(w, ParsedData); err != nil {
		log.Fatal("Failed to parse template ", err)
	}

	/*
		https://stackoverflow.com/questions/38190622/range-within-range-golang-template
		https://pkg.go.dev/html/template
	*/
}
