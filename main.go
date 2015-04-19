package main

import (
	"log"
	"mime"
	"net/http"
	"strings"

	"github.com/gorilla/pat"
	"github.com/ian-kent/go-angularjs-jquery-bootstrap-boilerplate/assets"
	"github.com/ian-kent/render"
)

type config struct {
	BindAddr string
}

var r *render.Render
var cfg config

func main() {
	cfg = config{
		BindAddr: ":3123",
	}

	r = render.New(render.Options{
		Asset:      assets.Asset,
		AssetNames: assets.AssetNames,
		Delims:     render.Delims{Left: "[:", Right: ":]"},
		Layout:     "layout",
	})

	pat := pat.New()
	pat.Path("/").Methods("GET").HandlerFunc(index)

	for _, file := range assets.AssetNames() {
		if strings.HasPrefix(file, "static/") {
			path := strings.TrimPrefix(file, "static")

			var mimeType string
			switch {
			case strings.HasSuffix(path, ".css"):
				mimeType = "text/css"
			case strings.HasSuffix(path, ".js"):
				mimeType = "application/javascript"
			default:
				mimeType = mime.TypeByExtension(path)
			}

			pat.Path(path).Methods("GET").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if b, err := assets.Asset("static" + path); err == nil {
					w.Header().Set("Content-Type", mimeType)
					w.WriteHeader(200)
					w.Write(b)
					return
				}
				w.WriteHeader(404)
			})
		}
	}

	err := http.ListenAndServe(cfg.BindAddr, pat)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "index", nil)
}
