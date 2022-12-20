package wdp

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Start() {
	url, _ := url.Parse("http://127.0.0.1:8080")
	proxy := httputil.NewSingleHostReverseProxy(url)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err == nil {
		fmt.Println(err)
	}
}
