package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/hostrouter"
	seed "github.com/jmulhern/seed/pkg"

	"github.com/jmulhern/whatever/pkg"
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
		packet := whatever.OpenPacket()
		if local := flags["local"]; local != "" {
			for _, l := range strings.Split(local, ",") {
				switch l {
				case "fqdn":
					for i, thing := range packet.Seeds {
						fqdnParts := strings.Split(thing.FQDN, ".")
						fqdnParts[len(fqdnParts)-1] = "local:3000"
						thing.FQDN = strings.Join(fqdnParts, ".")
						packet.Seeds[i] = thing
					}
				case "to":
					for i := range packet.Seeds {
						if len(packet.Seeds[i].Email.To) > 0 {
							packet.Seeds[i].Email.To = []string{"jmm@hey.com"}
						}
					}
				case "smtp":
					for i := range packet.Seeds {
						packet.Seeds[i].SMTP = seed.SMTP{}
					}
				case "bucket":
					for i := range packet.Seeds {
						packet.Seeds[i].Bucket = seed.Bucket{}
					}
				}
			}
		}

		if _, found := flags["peek"]; found {
			whatever.PeekAt(packet)
		}
		r := chi.NewRouter()
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.Timeout(60 * time.Second))

		hr := hostrouter.New()

		for _, seed := range packet.Seeds {
			switch seed.Name {
			case "whatever":
				h := whatever.NewHandler(seed)
				router := chi.NewRouter()
				router.Get("/dist/bundle.js", h.BundleJS)
				router.Get("/dist/bundle.css", h.BundleCSS)
				router.Get("/public/*", h.Public)
				router.Get("/*", h.Index)
				hr.Map(seed.FQDN, router)

			case "desert-cat-cookies":
				h := whatever.NewHandler(seed)
				router := chi.NewRouter()
				router.Get("/dist/bundle.js", h.BundleJS)
				router.Get("/dist/bundle.css", h.BundleCSS)
				router.Get("/public/*", h.Public)
				router.Post("/x/estimates", h.SubmitEstimate)
				router.Get("/*", h.Index)
				hr.Map(seed.FQDN, router)

			case "greasy-shadows", "the-bachelorette":
				h := whatever.NewHandler(seed)
				router := chi.NewRouter()
				router.Get("/dist/bundle.js", h.BundleJS)
				router.Get("/dist/bundle.css", h.BundleCSS)
				router.Get("/public/*", h.Public)
				router.Get("/*", h.Index)
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
