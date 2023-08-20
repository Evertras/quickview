package server

import (
	_ "embed"
	"os"

	"log"
	"net/http"
	"text/template"
)

//go:embed templates/index.html
var indexTemplateRaw string

var indexTemplate = template.Must(template.New("index").Parse(indexTemplateRaw))

func handlerIndex(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stat, err := os.Stat(filename)

		if err != nil {
			log.Println("ERROR FILE STAT:", err)
			w.WriteHeader(500)
			return
		}

		data := struct {
			Filename        string
			UnixNanoseconds int64
		}{
			Filename:        filename,
			UnixNanoseconds: stat.ModTime().UnixNano(),
		}

		err = indexTemplate.Execute(w, data)

		if err != nil {
			log.Println("ERROR RENDERING INDEX:", err)
		}
	}
}
