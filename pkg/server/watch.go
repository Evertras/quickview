package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Evertras/quickview/pkg/watcher"
	"golang.org/x/net/websocket"
)

func handlerWatcher(filename string) http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		log.Println("Opened connection")

		// Create a separate watcher for each instance, so that
		// we can freely pull from the channel
		fileWatcher, err := watcher.New(filename)

		if err != nil {
			log.Println("ERROR CREATING FILE WATCHER:", err)
			return
		}

		ctx, cancel := context.WithCancel(ws.Request().Context())
		defer cancel()

		go func() {
			buf := make([]byte, 64)
			for {
				// Read until we get an error, so we know our connection is closed...
				// the request context doesn't seem to close properly
				_, err := ws.Read(buf)
				if err != nil {
					cancel()
				}
			}
		}()

		timestamps, err := fileWatcher.Start(ctx)

		if err != nil {
			log.Println("ERROR STARTING FILE WATCHER:", err)
		}

		for nano := range timestamps {
			_, err = ws.Write([]byte(fmt.Sprintf("%v", nano)))

			if err != nil {
				log.Println("ERROR WRITING TO WEBSOCKET:", err)
				return
			}

			log.Println("Sent nano update:", nano)
		}

		log.Println("Request done")
	})
}
