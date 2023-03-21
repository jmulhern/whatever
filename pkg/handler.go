package whatever

import (
	"context"
	"crypto/tls"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-chi/chi/v5"
	seed "github.com/jmulhern/seed/pkg"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Handler struct {
	seed       seed.Seed
	awsConfig  aws.Config
	smtpServer *mail.SMTPServer
}

func NewHandler(seed seed.Seed) Handler {
	var smtpServer *mail.SMTPServer
	if seed.SMTP.Host != "" {
		smtpServer = mail.NewSMTPClient()
		smtpServer.Host = seed.SMTP.Host
		smtpServer.Port = seed.SMTP.Port
		smtpServer.Username = Decrypt(seed.SMTP.Username)
		smtpServer.Password = Decrypt(seed.SMTP.Password)
		smtpServer.Encryption = mail.EncryptionSTARTTLS
		smtpServer.KeepAlive = false
		smtpServer.ConnectTimeout = 10 * time.Second
		smtpServer.SendTimeout = 10 * time.Second
		smtpServer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	awsConfig, _ := config.LoadDefaultConfig(context.TODO())
	return Handler{
		seed:       seed,
		smtpServer: smtpServer,
		awsConfig:  awsConfig,
	}
}

func (h Handler) Index(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/index.html"))
	err := tmpl.Execute(w, map[string]any{
		"name": h.seed.Name,
		"site": h.seed.Site,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (Handler) Public(w http.ResponseWriter, r *http.Request) {
	rctx := chi.RouteContext(r.Context())
	pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir("public")))
	fs.ServeHTTP(w, r)
}

func (Handler) BundleJS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	raw, err := os.ReadFile("dist/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(raw)
}

func (Handler) BundleCSS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	raw, err := os.ReadFile("dist/bundle.css")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(raw)
}
