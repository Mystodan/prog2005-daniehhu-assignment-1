package handler

import (
	"net/http"
)

/*
Empty handler as default handler
*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No functionality on root level. Please use the path "+(RESOURCE_ROOT_PATH)+" for quick interface.", http.StatusOK)
}
