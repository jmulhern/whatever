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
	what "github.com/jmulhern/what/pkg"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Handler struct {
	thing      what.Thing
	awsConfig  aws.Config
	smtpServer *mail.SMTPServer
}

func NewHandler(thing what.Thing) Handler {
	var smtpServer *mail.SMTPServer
	if thing.SMTP.Host != "" {
		smtpServer = mail.NewSMTPClient()
		smtpServer.Host = thing.SMTP.Host
		smtpServer.Port = thing.SMTP.Port
		smtpServer.Username = thing.SMTP.Username
		smtpServer.Password = thing.SMTP.Password
		smtpServer.Encryption = mail.EncryptionSTARTTLS
		smtpServer.KeepAlive = false
		smtpServer.ConnectTimeout = 10 * time.Second
		smtpServer.SendTimeout = 10 * time.Second
		smtpServer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	awsConfig, _ := config.LoadDefaultConfig(context.TODO())
	return Handler{
		thing:      thing,
		smtpServer: smtpServer,
		awsConfig:  awsConfig,
	}
}

func (h Handler) Index(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/index.html"))
	err := tmpl.Execute(w, map[string]any{
		"Thing": h.thing,
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
