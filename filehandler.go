package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func basicHandler(rootpath string, w http.ResponseWriter, r *http.Request) {
	log.Println("rootpath", rootpath, "url", r.URL.Path)
	fpath := filepath.Join(rootpath, r.URL.Path)
	log.Println("fpath", fpath)

	if r.Method == "POST" {
		putHandler(fpath, w, r)
		return
	}

	fi, err := os.Stat(fpath)
	if err != nil {
		log.Printf("Couldn't stat %q because %v", fpath, err)
		// TODO(rjk): Figure out the right status code.
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if fi.IsDir() {
		getDirectory(rootpath, fpath, w)
		return
	}

	fd, err := os.Open(fpath)
	if err != nil {
		log.Println("Coudln't open", fpath, "because", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer fd.Close()
	setMimeType(fpath, w)
	if _, err := io.Copy(w, fd); err != nil {
		log.Println("failed to serve file", fpath, "because", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
