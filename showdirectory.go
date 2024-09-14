package main

import (
	//	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

// TODO(rjk): Make it prettier with nice CSS.
// TODO(rjk): Pull the CSS and JavaScript into a spearate payload (Go
// filesystem support.)

const dirlisting = `<html>
<head>
</head>
<body>
<ul id="filelist">
{{range .}} <li><a href="{{ .Path }}">{{ .Name }}</a></li> {{end}}
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

// event is a drop event (the ondrop event handler.)
function dodrop(event)
{
	let dt = event.dataTransfer;
	let files = dt.files;
 
	([...files]).forEach(uploadFile)
}


function uploadFile(file) {
	console.log(file)

	let url = window.location.href;
	let formData = new FormData();

	formData.append('file', file);

  fetch(url, {
    method: 'POST',
    body: formData
  })
  .then(() => { /* Done. Inform the user */ 
	console.log("have done the upload");
	let filelist = document.getElementById("filelist")
	filelist.insertAdjacentHTML("beforeend", '<li> <a href=./"' + file.name + '">' + file.name + '</a></li>');
	})
  .catch(() => { /* Error. Inform the user */ 
	console.log("didn't work out, the upload thing. shucks");
	// TODO(rjk): Do something pretty here.
	})
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

type FileStructure struct {
	Name string
	Path string
}

func getDirectory(rootpath, fpath string, w http.ResponseWriter) {
	log.Println("should do something here for a directory", fpath)

	// TODO(rjk): I canz not parse this on every request? But hey, it doesn't
	// matter much?
	t := template.Must(template.New("dirlisting").Parse(dirlisting))

	direntries, err := os.ReadDir(fpath)
	if err != nil {
		log.Printf("Couldn't read dir %q because %v", fpath, err)
		w.WriteHeader(http.StatusNotFound)
	}

	fss := make([]FileStructure, 0)
	for _, de := range direntries {
		ap := path.Join(fpath, de.Name())
		fss = append(fss, FileStructure{Name: de.Name(), Path: ap[len(rootpath):]})
	}

	if err := t.Execute(w, fss); err != nil {
		log.Println("Can't run template?", err)
	}
}
