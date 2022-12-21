package wdp

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/radovskyb/watcher"
)

var w *watcher.Watcher

func Watcher() {
	w = watcher.New()
	go watch()

	if err := w.AddRecursive(Watch); err != nil {
		color.Red(err.Error())
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		color.Red(err.Error())
	}
}

func watch() {
	for {
		event := <-w.Event
		color.Yellow(fmt.Sprintf("File changed: %s", event.Path))
		for _, socket := range sockets {
			err := socket.WriteMessage(websocket.BinaryMessage, []byte{})
			if err != nil {
				color.Red(err.Error())
			}
		}
	}
}
