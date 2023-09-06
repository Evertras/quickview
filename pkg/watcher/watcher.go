package watcher

import (
	"context"
	"fmt"
	"log"
	"path"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	fsWatcher *fsnotify.Watcher
	filename  string
	stop      func()

	nanoTimestamps chan int64
}

func New(filename string) (*Watcher, error) {
	fsWatcher, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, fmt.Errorf("failed to create fs watcher: %w", err)
	}

	return &Watcher{
		filename:       filename,
		fsWatcher:      fsWatcher,
		nanoTimestamps: make(chan int64),
	}, nil
}

func (w *Watcher) Start(ctx context.Context) (<-chan int64, error) {
	if w.stop != nil {
		return nil, fmt.Errorf("already started")
	}

	ctx, cancel := context.WithCancel(ctx)

	w.stop = cancel

	watchPath := path.Dir(w.filename)
	log.Println("Watching", watchPath)
	err := w.fsWatcher.Add(watchPath)
	if err != nil {
		return nil, fmt.Errorf("failed to add watch path %q: %w", watchPath, err)
	}

	go func() {
		var lastNano int64 = 0

		defer func() {
			close(w.nanoTimestamps)

			for err := range w.fsWatcher.Errors {
				log.Println("ERROR FROM FSWATCHER:", err)
			}

			err := w.fsWatcher.Close()
			if err != nil {
				log.Println("ERROR: failed to close internal fsWatcher:", err)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return

			case err := <-w.fsWatcher.Errors:
				log.Println("ERROR FROM FSWATCHER:", err)

			case event, ok := <-w.fsWatcher.Events:
				if !ok {
					log.Println("Internal fswatcher closed")
					return
				}

				if event.Name != w.filename {
					continue
				}

				nano, err := getFileModifiedNano(w.filename)

				if err != nil {
					log.Println("ERROR GETTING MODIFIED TIME:", err)
					return
				}

				if nano == lastNano {
					continue
				}

				lastNano = nano

				w.nanoTimestamps <- nano
			}
		}
	}()

	return w.nanoTimestamps, nil
}
