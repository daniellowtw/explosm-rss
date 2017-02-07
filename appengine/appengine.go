package main

import (
	"net/http"
	"time"

	"github.com/daniellowtw/explosm-rss/explosm"
)

var timeoutDuration = time.Hour

func init() {
	e := explosm.Explosm{
		RefreshInterval: time.Hour,
	}
	var lastUpdated time.Time
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is a hack because it seems like I can't do polling in the background without a user generated request
		if time.Now().Sub(lastUpdated) > timeoutDuration {
			e.Do(r)
			lastUpdated = time.Now()
		}
		w.Header().Add("content-type", "text/xml")
		w.Write([]byte(e.Generate()))
	})
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "image/png")
		w.Write(explosm.Icon96)
	})
}
