package whatever

import (
	"log"
	"net/http"
	"os"
)

func (Handler) GetBundleJS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=31536000")
	raw, err := os.ReadFile("dist/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(raw)
}

func (Handler) GetBundleCSS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "max-age=31536000")
	raw, err := os.ReadFile("dist/bundle.css")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(raw)
}
