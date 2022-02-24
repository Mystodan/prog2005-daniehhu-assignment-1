package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getWebStatus(inn string) string {
	stats, err := http.Get(inn)
	checkError(err)
	return stats.Status
}

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	msg := Diagnostics{
		getWebStatus(GET_UNI),
		getWebStatus(GET_CNTR),
		VERSION,
		fmt.Sprintf("%f", getUptime(timer).Seconds()) + "s",
	}
	w.Header().Add("content-type", "application/json")

	// Instantiate encoder
	encoder := json.NewEncoder(w)
	// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
	err := encoder.Encode(msg)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}
}
