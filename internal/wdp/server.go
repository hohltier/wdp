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
	"github.com/sho7a/wdp/web"
)

func Start() {
	url, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", Port))
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ModifyResponse = modify

	fs, _ := fs.Sub(web.Static, "static")
	http.Handle("/.wdp/", http.StripPrefix("/.wdp/", http.FileServer(http.FS(fs))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", Listen))
	fmt.Println(fmt.Sprintf("Listening on :%d", listener.Addr().(*net.TCPAddr).Port))
	if err := http.Serve(listener, nil); err != nil {
		fmt.Println(err)
	}
}

func modify(r *http.Response) error {
	if r.StatusCode == http.StatusOK && strings.HasPrefix(r.Header.Get("Content-Type"), "text/html") {
		defer r.Body.Close()
		doc, _ := goquery.NewDocumentFromReader(r.Body)
		doc.Find("head").AppendHtml(`<script src="/.wdp/js/wdp.js"></script>`)
		h, _ := doc.Html()
		b := []byte(h)
		r.Body = io.NopCloser(bytes.NewReader(b))
		r.ContentLength = int64(len(b))
		r.Header.Set("Content-Length", strconv.Itoa(len(b)))
	}
	return nil
}
