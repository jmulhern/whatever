package whatever

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/jsii-runtime-go"
	mail "github.com/xhit/go-simple-mail/v2"
	"gopkg.in/yaml.v3"

	"github.com/jmulhern/whatever/pkg/kind"
)

func (h Handler) SubmitEstimate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	quantity := -1
	if parsed, err := strconv.Atoi(strings.TrimSpace(r.Form.Get("quantity"))); err == nil {
		quantity = parsed
	} else {
		log.Fatal(err)
	}

	estimate := kind.Estimate{
		FirstName:   strings.TrimSpace(r.Form.Get("first-name")),
		LastName:    strings.TrimSpace(r.Form.Get("last-name")),
		Email:       strings.TrimSpace(r.Form.Get("email")),
		PhoneNumber: strings.TrimSpace(r.Form.Get("phone-number")),
		Theme:       strings.TrimSpace(r.Form.Get("theme")),
		Quantity:    quantity,
		NeededBy:    strings.TrimSpace(r.Form.Get("needed-by")),
		Message:     strings.TrimSpace(r.Form.Get("message")),
	}

	fmt.Println("--- estimate start ---")
	raw, _ := json.MarshalIndent(estimate, "", "  ")
	fmt.Println(string(raw))
	fmt.Println("--- estimate stop ---")

	if h.seed.Bucket.Name != "" && estimate.HasContactInfo() {
		fmt.Println("--- bucket start ---")
		now := time.Now()

		name := "unknown"
		if estimate.FirstName != "" && estimate.LastName != "" {
			name = strings.ToLower(fmt.Sprintf("%s-%s", estimate.LastName, estimate.FirstName))
		}

		raw, _ := yaml.Marshal(estimate)
		key := fmt.Sprintf("estimates/%d/%d/%d/%s-%s.yaml",
			now.Year(), now.Month(), now.Day(), now.Format("150405"), name)

		_, err = s3.NewFromConfig(h.awsConfig).PutObject(r.Context(), &s3.PutObjectInput{
			Bucket: jsii.String(h.seed.Bucket.Name),
			Key:    jsii.String(key),
			Body:   bytes.NewReader(raw),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created s3://%s/%s\n", h.seed.Bucket.Name, key)
		fmt.Println("--- bucket stop ---")
	}

	if h.seed.SMTP.Host != "" && estimate.HasContactInfo() {
		fmt.Println("--- email start ---")
		raw, _ := yaml.Marshal(estimate)
		body := fmt.Sprintf(`
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello</title>
</head>
<body>
<pre>%s</pre>
</body>
`, string(raw))

		email := mail.NewMSG()
		email.SetFrom(h.seed.Email.From)
		email.AddTo(h.seed.Email.To...)
		email.SetSubject(fmt.Sprintf("Estimate Request - %s %s", estimate.FirstName, estimate.LastName))
		email.SetBody(mail.TextHTML, body)

		smtpClient, err := h.smtpServer.Connect()
		if err != nil {
			log.Fatal(err)
		}

		err = email.Send(smtpClient)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Email Sent for ", estimate.FirstName, estimate.LastName)
		}

		if email.Error != nil {
			log.Fatal(email.Error)
		}
		fmt.Println("--- email stop ---")
	}
	http.Redirect(w, r, "/?submitted=true", http.StatusFound)
}
