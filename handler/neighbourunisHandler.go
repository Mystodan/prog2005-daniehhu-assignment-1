package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

/*
Entry point handler for Location information
*/
func NBuinfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		NBGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET or POST are supported.", http.StatusNotImplemented)
		return
	}
}

/*
Dedicated handler for GET requests
*/
func NBGetRequest(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/")
	var urlWant int
	errCode := false
	comp := strings.ReplaceAll(NEIGHBOURUNIS_PATH, "/", "")
	//fmt.Println(urlSplit, ":", urlSplit[len(urlSplit)-1], ":", len(urlSplit))
	for i, s := range urlSplit {
		if s == (comp) {
			urlWant = i + 1
		}
	}
	var firstAppendVal, secondAppendVal string
	firstAppendVal = strings.ReplaceAll(urlSplit[urlWant], " ", "%20")
	if len(firstAppendVal) > 0 {
		firstAppendVal = "=" + firstAppendVal
	} else {
		errCode = true
	}
	if len(urlSplit)-1 == urlWant+1 {
		secondAppendVal = strings.ReplaceAll(urlSplit[urlWant+1], " ", "%20")
		if !(len(secondAppendVal) > 0) {
			errCode = true
		}
	} else {
		errCode = true
	}

	if errCode {
		fmt.Println(secondAppendVal, firstAppendVal)
		http.Error(w, "No functionality without parameters: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}", http.StatusOK)
	} else {

		fmt.Println(secondAppendVal, firstAppendVal)
		write := getURL(GET_UNI + UNI_REQ + secondAppendVal + "&country" + firstAppendVal)

		var getLimit int64
		getParam := strings.Split(r.URL.RawQuery, "limit=")
		if len(getParam) > 1 {
			t, _ := strconv.ParseInt(getParam[1], 10, 0)
			getLimit = t
		}

		var getU []getUnii
		body, err := io.ReadAll(write.Body)

		checkError(err)
		json.Unmarshal(body, &getU)
		// Write content type header (best practice)
		w.Header().Add("content-type", "application/json")

		// Instantiate encoder
		encoder := json.NewEncoder(w)
		var setUni []Universities
		setUni = append(setUni, setUniversity(getU)...)
		setUni = append(setUni, getBorderingUniversities(setUni, int(getLimit))...)

		// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
		err = encoder.Encode(setUni)
		checkError(err)
	}
}
