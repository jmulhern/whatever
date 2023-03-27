package whatever

import (
	"log"
	"math/rand"
	"net/http"
)

func (h Handler) GetIndex(w http.ResponseWriter, _ *http.Request) {
	nonce := makeNonce()
	if csp := h.seed.ContentSecurityPolicy; csp.Make(nonce) != "" {
		if csp.ReportOnly {
			w.Header().Set("Content-Security-Policy-Report-Only", csp.Make(nonce))
		} else {
			w.Header().Set("Content-Security-Policy", csp.Make(nonce))
		}
	}

	err := h.index.Execute(w, map[string]any{
		"name":  h.seed.Name,
		"site":  h.seed.Site,
		"nonce": nonce,
		"cdn": h.seed.CDN.Use,
	})
	if err != nil {
		log.Fatal(err)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func makeNonce() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
