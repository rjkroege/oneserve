package main

import (
	//	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// TODO(rjk): Make it prettier with nice CSS.

const dirlisting = `<html>
<body>
<ul>
{{range .}} <li><a href="./{{ .Name }}">{{ .Name }}</a></li> {{end}}
</ul>
</body>
</html>
`

func getDirectory(fpath string, w http.ResponseWriter) {
	log.Println("should do something here for a directory")

	// TODO(rjk): I canz not parse this on every request? But hey, it doesn't
	// matter much?
	t := template.Must(template.New("dirlisting").Parse(dirlisting))

	direntries, err := os.ReadDir(fpath)
	if err != nil {
		log.Printf("Couldn't read dir %q because %v", fpath, err)
		w.WriteHeader(http.StatusNotFound)
	}
	log.Println("read a directory")

	if err := t.Execute(w, direntries); err != nil {
		log.Println("Can't run template?", err)
	}
}
