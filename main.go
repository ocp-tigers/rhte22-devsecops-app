package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const version = "1.2"

func healthEndpoint(w http.ResponseWriter, r *http.Request) {
	health := "Healthy"
	_, err := w.Write([]byte(health + "\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func serveApp(address string, infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Printf("Starting Server on %s. Version %s", address, version)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/health", healthEndpoint)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		errorLog.Fatal(err)
	}
}

func main() {
	// All incoming HTTP requests are server in their own goroutine
	addr := flag.String("addr", ":8080", "HTTP network address")

	// Parse the command line and assigns it to the addr variable
	flag.Parse()
	// Use log.New() to create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// writing error messages using the stderr as the output
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	serveApp(*addr, infoLog, errorLog)
}
