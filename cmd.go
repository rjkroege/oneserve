package main

import (
	"io"
	"log"
	"flag"
	"net/http"
	"os"
	"strconv"
)


var portFlag  int

func init() {
	flag.IntVar(&portFlag, "port", 8000, "port number")
}

func main() {
	log.Println("Hello")

	flag.Parse();

	args := flag.Args()

	for _, fn := range args {
		log.Println("file to serve", fn)

		http.HandleFunc("/" + fn, func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			relpth := r.URL.Path[1:]
			
			fd, err  := os.Open(relpth)
			if err != nil {
				// should fail out in some way?
				log.Println("Coudln't open", relpth, "because", err)
				// write something...
				
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if _, err := io.Copy(w, fd); err != nil {
				log.Println("failed to serve file",  relpth, "because", err)
				w.WriteHeader(http.StatusBadRequest)
			}
		})
	}

	network := ":" + strconv.Itoa(portFlag)
	log.Fatal(http.ListenAndServe(network, nil))


}