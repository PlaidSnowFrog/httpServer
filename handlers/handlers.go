package handlers

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to the homepage, go to /help for more information")
}

func Help(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to the help page\n
				Go to /send to send data that will be saved") // still not finished
}

func Save(w http.ResponseWriter, r *http.Request, filePath string) {
	
}
