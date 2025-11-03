package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"time"

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
	var strapiRes models.EventsPageData

	client := &http.Client{Timeout: time.Second * 10}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("There was no response from Strapi")
			}
		}()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read response")
	} else {
		json.Unmarshal(body, &strapiRes)
		fmt.Println(strapiRes)
	}

	defer res.Body.Close()

	return strapiRes
}

func FetchTribeMembers(url string, token string) models.TribeMembers {
	godotenv.Load()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result models.TribeMembers
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Unable to marshal JSON")
	}
	return result
}

func FetchMembershipDocs(url string, token string) models.ProgramDocs {
	godotenv.Load()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result models.ProgramDocs
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Unable to marshal JSON")
	}

	return result
}

func FetchProgramDocs(url string, token string) models.ProgramDocs {
	godotenv.Load()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result models.ProgramDocs
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
	token := os.Getenv("TRIBE_MEMBERS_BEARER_TOKEN")
	url := os.Getenv("ABOUT_ENDPOINT")
	ParsedData := FetchTribeMembers(url, token).Data

	if err := checkPath("/about", w, r); err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("static/templates/about.html")

	if err := tmpl.Execute(w, ParsedData); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("EVENTS_BEARER_TOKEN")
	url := os.Getenv("EVENTS_ENDPOINT")
	ParsedData := FetchEvents(url, token)

	if err := checkPath("/events", w, r); err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("static/templates/events.html")

	if err := tmpl.Execute(w, ParsedData); err != nil {
		log.Fatal("Failed to parse template 181 ", err)
	}

}

func Membership(w http.ResponseWriter, r *http.Request) {}
func Programs(w http.ResponseWriter, r *http.Request)   {}
