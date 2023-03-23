package whatever

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (Handler) GetPublic(w http.ResponseWriter, r *http.Request) {
	rctx := chi.RouteContext(r.Context())
	pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir("public")))
	w.Header().Set("Cache-Control", "max-age=31536000")
	fs.ServeHTTP(w, r)
}
