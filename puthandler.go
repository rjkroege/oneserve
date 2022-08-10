package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func putHandler(rootpath string, w http.ResponseWriter, r *http.Request) {
	// log.Println("should handle the put now")
	// log.Println(r)

	mpf, mpinfo, err := r.FormFile("file")
	if err != nil {
		log.Println("can't parse the payload:", err)
		return
	}
	defer mpf.Close()

	destfilepath := filepath.Join(rootpath, mpinfo.Filename)
	dest, err := os.Create(destfilepath)
	if err != nil {
		log.Printf("can't open %q: %v]n", destfilepath, err)
		return
	}
	defer dest.Close()
	if _, err := io.Copy(dest, mpf); err != nil {
		log.Printf("can't copy POST contents to %q: %v\n", destfilepath, err)
	}
}
