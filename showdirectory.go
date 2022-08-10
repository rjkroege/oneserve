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
<head>
</head>
<body>
<ul>
{{range .}} <li><a href="./{{ .Name }}">{{ .Name }}</a></li> {{end}}
</ul>
<br>
<div id="output" style="min-height: 200px; white-space: pre; border: 1px solid black;"
     ondragenter="document.getElementById('output').textContent = ''; event.stopPropagation(); event.preventDefault();"
     ondragover="event.stopPropagation(); event.preventDefault();"
     ondrop="event.stopPropagation(); event.preventDefault();
     dodrop(event);">
     Drop files to upload here!
</div>
<script>

function dodrop(event)
{
  var dt = event.dataTransfer;
  var files = dt.files;

  var count = files.length;
  output("File Count: " + count + "\n");

    for (var i = 0; i < files.length; i++) {
      output(" File " + i + ":\n(" + (typeof files[i]) + ") : <" + files[i] + " > " +
             files[i].name + " " + files[i].size + "\n");
    }
}

function output(text)
{
  document.getElementById("output").textContent += text;
  //dump(text);
}
</script>
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
