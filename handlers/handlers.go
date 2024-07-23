package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func Home(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to the homepage, go to /help for more information")
}

func Help(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to the help page\nGo to /send to send data that will be saved")
}

func Save(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.OpenFile("logged.rbgc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Serialize the request to a byte slice
	dumpedReq, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		log.Printf("Failed to dump request: %v", err)
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	ip := string(dumpedReq)

	time := time.Now().Format("2006-01-02 15:04:05") // YYYY-MM-DD HH-MM-SS

	var text string = time + ip + "\n"

	if _, err := file.WriteString(text); err != nil {
		log.Fatal(err)
	}
}
