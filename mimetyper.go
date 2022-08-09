package main

import (
	"net/http"
	"path"
)

// I copied this from blog. It behooves me to create a library that
// I can use.
func setMimeType(pth string, w http.ResponseWriter) {
	ext := path.Ext(pth)

	// Expand this as necessary based on the content that we feature.
	switch ext {
	case "":
		pth = path.Join(pth, "index.html")
		w.Header().Set("Content-Type", "text/html")
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	case ".xml":
		w.Header().Set("Content-Type", "application/xml")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case ".jpeg", ".jpg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	}
}
