package whatever

import (
	"context"
	"crypto/tls"
	"embed"
	"html/template"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	heritage "github.com/jmulhern/heritage/pkg"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Handler struct {
	seed       heritage.Seed
	awsConfig  aws.Config
	smtpServer *mail.SMTPServer
	index      *template.Template
}

func NewHandler(fs embed.FS, seed heritage.Seed) Handler {
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
	index := template.Must(template.ParseFS(fs, "templates/*.html"))
	return Handler{
		seed:       seed,
		smtpServer: smtpServer,
		awsConfig:  awsConfig,
		index:      index,
	}
}
