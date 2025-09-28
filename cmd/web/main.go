package main

import (
	"fmt"
	handlers "github/strapi-admin/pkg/handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	const PORT = 3000
	godotenv.Load()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Listening on port %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
