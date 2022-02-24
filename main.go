package main

import (
	"assignment-1/handler"
	"log"
	"net/http"
	"os"
)

func handleRequests(port string) {

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func setPort(inn string) {
	os.Setenv("PORT", inn)
}

func getPort() string {
	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}
	return port
}

func main() {
	setPort("") // default port is :8080
	handler.TimerStart()
	// Set up handler endpoints
	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.UNIINFO_PATH, handler.UniinfoHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.NEIGHBOURUNIS_PATH, handler.NBuinfoHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.DIAG_PATH, handler.DiagHandler)

	handleRequests(getPort())
}
