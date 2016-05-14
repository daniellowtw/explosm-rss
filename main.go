package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/daniellowtw/explosm-rss/explosm"
)

func main() {
	var port = flag.Int64("port", 20480, "port to run the server on")
	var refreshInterval = flag.Duration("refresh_interval", time.Hour*1, "interval to check for new data")

	flag.Parse()
	e := explosm.Explosm{
		RefreshInterval: *refreshInterval,
	}
	a := make(chan struct{})
	go e.Run(a)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/xml")
		w.Write([]byte(e.Generate()))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}