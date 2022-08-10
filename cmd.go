package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var addrFlag string

func init() {
	flag.StringVar(&addrFlag, "addr", "localhost:8000", "address")
	flag.Usage = usage
}

const usagestring = `oneserve [<file>...]
Provide a list of one or more files to seve those and only those
where the root of the http tree is the cwd where oneserve runs.
In the absence of arguments, serves all of the files in the cwd.

Paths must be w.r.t. the cwd.
`

func usage() {
	os.Stdout.WriteString(usagestring)
	flag.PrintDefaults()
}

func main() {
	log.Println("Hello")

	flag.Parse()
	args := flag.Args()
	rootpath, err := os.Getwd()
	if err != nil {
		log.Fatalln("can't determine current directory: ", err)
	}

	for _, fn := range args {
		log.Println("file to serve", fn)

		bfn, err := filepath.Abs(fn)
		if err != nil {
			log.Fatalf("Can't absolute %s: %v\n", fn, err)
		}
		relfn, err := filepath.Rel(rootpath, bfn)
		if err != nil {
			log.Fatalf("Bad path %s: %v\n", fn, err)
		}
		if len(relfn) > 1 && relfn[0:2] == ".." {
			log.Fatalf("%s not allowed to climb out of %s\n", relfn, rootpath)
		}

		http.HandleFunc(path.Join("/", relfn), func(w http.ResponseWriter, r *http.Request) {
			basicHandler(rootpath, w, r)
		})
	}

	if len(args) == 0 {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			fpath := filepath.Join(rootpath, r.URL.Path)

			fd, err := os.Open(fpath)
			if err != nil {
				log.Println("Coudln't open", fpath, "because", err)
				w.WriteHeader(http.StatusNotFound)
				return
			}
			setMimeType(fpath, w)
			if _, err := io.Copy(w, fd); err != nil {
				log.Println("failed to serve file", fpath, "because", err)
				w.WriteHeader(http.StatusBadRequest)
			}
		})
	}

	log.Fatal(http.ListenAndServe(addrFlag, nil))
}
