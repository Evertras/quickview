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

type indexData struct {
	filename     string
	websocketURL string
}

func handlerIndex(data indexData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stat, err := os.Stat(data.filename)

		if err != nil {
			log.Println("ERROR FILE STAT:", err)
			w.WriteHeader(500)
			return
		}

		d := struct {
			Filename        string
			UnixNanoseconds int64
			WebsocketURL    string
		}{
			Filename:        data.filename,
			UnixNanoseconds: stat.ModTime().UnixNano(),
			WebsocketURL:    data.websocketURL,
		}

		err = indexTemplate.Execute(w, d)

		if err != nil {
			log.Println("ERROR RENDERING INDEX:", err)
		}
	}
}
