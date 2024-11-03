// handlers/staticserver.go
package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"tracker/api"
)

// StaticServer serves static files in the /static/ directory
func StaticServer(w http.ResponseWriter, r *http.Request) {
	// Check if the path starts with "/static/"
	if r.URL.Path == "/static/" {
		api.HandleError(w, nil, http.StatusForbidden, "403.html")
		return
	}

	// Remove "/static" prefix and join with the static file directory
	filePath := filepath.Join("static", r.URL.Path[len("/static/"):])

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		api.HandleError(w, nil, http.StatusNotFound, "404.html")
		return
	}
	if err != nil {
		api.HandleError(w, nil, http.StatusInternalServerError, "500.html")
		return
	}
	if info.IsDir() {
		api.HandleError(w, nil, http.StatusForbidden, "403.html")
		return
	}

	http.ServeFile(w, r, filePath)
}
