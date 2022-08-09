package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getDirectory(fpath string, w http.ResponseWriter) {
	log.Println("should do something here for a directory")

	direntries, err := os.ReadDir(fpath)
	if err != nil {
		log.Printf("Couldn't read dir %q because %v", fpath, err)
		w.WriteHeader(http.StatusNotFound)
	}
	log.Println("read a directory")

	fmt.Fprintf(w, "<html><body><ul>\n")
	w.Header().Set("Content-Type", "text/html")
	for _, de := range direntries {
		log.Println(de)
		fmt.Fprintf(w, "<li>%s</li>\n", de.Name())
	}
	fmt.Fprintf(w, "</ul></body></html>\n")
}
