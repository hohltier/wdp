package wdp

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/sho7a/wdp/web"
	"golang.org/x/exp/slices"
)

var (
	sockets  []*websocket.Conn
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func Server() {
	url, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", Port))
	if err != nil {
		color.Red(err.Error())
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ModifyResponse = modify

	http.HandleFunc("/.wdp/ws", socket)

	fs, err := fs.Sub(web.Static, "static")
	if err != nil {
		color.Red(err.Error())
	}
	http.Handle("/.wdp/", http.StripPrefix("/.wdp/", http.FileServer(http.FS(fs))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Listen))
	if err != nil {
		color.Red(err.Error())
	}
	color.Green(fmt.Sprintf("Listening on :%d", listener.Addr().(*net.TCPAddr).Port))
	if err := http.Serve(listener, nil); err != nil {
		color.Red(err.Error())
	}
}

func modify(r *http.Response) error {
	if r.StatusCode == http.StatusOK && strings.HasPrefix(r.Header.Get("Content-Type"), "text/html") {
		defer r.Body.Close()
		doc, err := goquery.NewDocumentFromReader(r.Body)
		if err != nil {
			color.Red(err.Error())
			return err
		}
		doc.Find("head").AppendHtml(`<script src="/.wdp/js/wdp.js"></script>`)
		h, err := doc.Html()
		if err != nil {
			color.Red(err.Error())
			return err
		}
		b := []byte(h)
		r.Body = io.NopCloser(bytes.NewReader(b))
		r.ContentLength = int64(len(b))
		r.Header.Set("Content-Length", strconv.Itoa(len(b)))
	}
	return nil
}

func socket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		color.Red(err.Error())
	}
	sockets = append(sockets, ws)
	_, _, err = ws.ReadMessage()
	if err != nil {
		i := slices.Index(sockets, ws)
		sockets = append(sockets[:i], sockets[i+1:]...)
	}
}
