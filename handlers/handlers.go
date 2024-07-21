package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage, go to /help for more information")
}

func Help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the help page\n
				Go to /send to send data that will be saved") // still not finished
}
