package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/hostrouter"
	heritage "github.com/jmulhern/heritage/pkg"

	"github.com/jmulhern/whatever/pkg"
)

var (
	//go:embed templates
	templates embed.FS
)

func main() {
	if len(os.Args) <= 1 {
		panic("too few arguments")
	}
	action := os.Args[1]
	flags := make(map[string]string)
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			argParts := strings.Split(arg, "=")
			key := strings.TrimPrefix(argParts[0], "--")

			var value string
			if len(argParts) > 1 {
				value = strings.Join(argParts[1:], "=")
			} else {
				value = "true"
			}
			flags[key] = value
		}
	}

	fmt.Println(action, flags)
	switch action {
	case "encrypt":
		fmt.Println(whatever.Encrypt(os.Args[2]))
	case "decrypt":
		fmt.Println(whatever.Decrypt(os.Args[2]))
	case "serve":
		var packet heritage.Packet
		if local := flags["local"]; local == "" {
			packet = whatever.OpenPacket()
		} else {
			packet = whatever.OpenLocalPacket(strings.Split(local, ",")...)
		}
		if _, found := flags["peek"]; found {
			whatever.PeekAt(packet)
		}
		r := chi.NewRouter()
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.Compress(5))
		r.Use(middleware.Timeout(60 * time.Second))

		hr := hostrouter.New()

		for _, seed := range packet.Seeds {
			switch seed.Name {
			// Live
			case "whatever":
				h := whatever.NewHandler(templates, seed)
				router := chi.NewRouter()
				// backend
				router.Get("/x/things", h.GetThings)
				// frontend
				router.Get("/dist/bundle.js", h.GetBundleJS)
				router.Get("/dist/bundle.css", h.GetBundleCSS)
				router.Get("/public/*", h.GetPublic)
				router.Get("/", h.GetIndex)
				// support
				router.Post("/report/csp", h.ReceiveContentSecurityPolicyReport)
				router.Post("/health", h.GetHealth)
				// not found
				router.Get("/*", h.Get404)
				hr.Map(seed.FQDN, router)

			case "desert-cat-cookies":
				h := whatever.NewHandler(templates, seed)
				router := chi.NewRouter()
				// backend
				router.Post("/x/estimates", h.CreateEstimate)
				// frontend
				router.Get("/dist/bundle.js", h.GetBundleJS)
				router.Get("/dist/bundle.css", h.GetBundleCSS)
				router.Get("/public/*", h.GetPublic)
				router.Get("/", h.GetIndex)
				// support
				router.Post("/report/csp", h.ReceiveContentSecurityPolicyReport)
				router.Post("/health", h.GetHealth)
				// not found
				router.Get("/*", h.Get404)
				hr.Map(seed.FQDN, router)

			case "greasy-shadows":
				h := whatever.NewHandler(templates, seed)
				router := chi.NewRouter()
				// frontend
				router.Get("/dist/bundle.js", h.GetBundleJS)
				router.Get("/dist/bundle.css", h.GetBundleCSS)
				router.Get("/public/*", h.GetPublic)
				router.Get("/", h.GetIndex)
				// support
				router.Post("/report/csp", h.ReceiveContentSecurityPolicyReport)
				router.Post("/health", h.GetHealth)
				// not found
				router.Get("/*", h.Get404)
				hr.Map(seed.FQDN, router)

			case "watch":
				h := whatever.NewHandler(templates, seed)
				router := chi.NewRouter()
				// backend
				router.Get("/x/movies", h.GetMovies)
				// frontend
				router.Get("/dist/bundle.js", h.GetBundleJS)
				router.Get("/dist/bundle.css", h.GetBundleCSS)
				router.Get("/public/*", h.GetPublic)
				router.Get("/", h.GetIndex)
				// support
				router.Post("/report/csp", h.ReceiveContentSecurityPolicyReport)
				router.Post("/health", h.GetHealth)
				// default
				router.Get("/*", h.Get404)
				hr.Map(seed.FQDN, router)

			// Incubator
			case "the-bachelorette":
				h := whatever.NewHandler(templates, seed)
				router := chi.NewRouter()
				// frontend
				router.Get("/dist/bundle.js", h.GetBundleJS)
				router.Get("/dist/bundle.css", h.GetBundleCSS)
				router.Get("/public/*", h.GetPublic)
				router.Get("/", h.GetIndex)
				// default
				router.Get("/*", h.Get404)
				hr.Map(seed.FQDN, router)

			case "hall-of-fame":
				h := whatever.NewHandler(templates, seed)
				router := chi.NewRouter()
				// frontend
				router.Get("/dist/bundle.js", h.GetBundleJS)
				router.Get("/dist/bundle.css", h.GetBundleCSS)
				router.Get("/public/*", h.GetPublic)
				router.Get("/", h.GetIndex)
				// default
				router.Get("/*", h.Get404)
				hr.Map(seed.FQDN, router)
			}
		}

		// everything else
		routes := chi.NewRouter()
		routes.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("OK"))
		})
		routes.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("..."))
		})
		hr.Map("*", routes)
		r.Mount("/", hr)
		log.Println(http.ListenAndServe(":3000", r))
	}
}
