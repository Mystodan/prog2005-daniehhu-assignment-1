package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

/*
Entry point handler for Location information
*/
func UniinfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		UniGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET or POST are supported.", http.StatusNotImplemented)
		return
	}
}

/*
Dedicated handler for GET requests
*/
func UniGetRequest(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/")
	var urlWant int
	comp := strings.ReplaceAll(UNIINFO_PATH, "/", "")
	for i, s := range urlSplit {
		if s == (comp) {
			urlWant = i + 1
		}
	}

	lastAppendVal := strings.ReplaceAll(urlSplit[urlWant], " ", "%20")
	if len(lastAppendVal) > 0 {
		write := getURL(GET_UNI + UNI_REQ + lastAppendVal)

		var getU []getUnii
		body, err := io.ReadAll(write.Body)

		checkError(err)
		json.Unmarshal(body, &getU)
		// Write content type header (best practice)
		w.Header().Add("content-type", "application/json")

		// Instantiate encoder
		encoder := json.NewEncoder(w)

		// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
		err = encoder.Encode(setUniversity(getU))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "No functionality without parameters: uniinfo/{:partial_or_complete_university_name}/", http.StatusOK)
	}

}
