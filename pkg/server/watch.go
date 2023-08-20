package server

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/net/websocket"
)

func handlerWatcher(filename string) http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		log.Println("Opened connection")

		// Create a separate watcher for each instance, so that
		// we can freely pull from the channel
		watcher, err := fsnotify.NewWatcher()

		if err != nil {
			log.Println("ERROR WATCHING FILE:", err)
			return
		}

		watchPath := path.Dir(filename)
		log.Println("Watching", watchPath)
		err = watcher.Add(watchPath)
		if err != nil {
			log.Println("ERROR ADDING FILE WATCH:", err)
		}
		defer watcher.Close()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		var lastNano int64 = 0

		for event := range watcher.Events {
			if event.Name != filename {
				continue
			}

			nano, err := getFileModifiedNano(filename)

			if err != nil {
				log.Println("ERROR GETTING MODIFIED TIME:", err)
				return
			}

			if nano == lastNano {
				continue
			}

			lastNano = nano

			_, err = ws.Write([]byte(fmt.Sprintf("%v", nano)))

			if err != nil {
				log.Println("ERROR WRITING TO WEBSOCKET:", err)
				return
			}

			log.Println("Sent nano update:", nano)
		}

		for err := range watcher.Errors {
			log.Println("ERROR FROM WATCHER:", err)
		}
	})
}
