package whatever

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/jsii-runtime-go"
	"gopkg.in/yaml.v3"
)

type ContentSecurityPolicyRequest struct {
	CSPReport ContentSecurityPolicyReport `json:"csp-report"`
}

type ContentSecurityPolicyReport struct {
	DocumentURI        string `json:"document-uri" yaml:"document-uri,omitempty"`
	Referrer           string `json:"referrer" yaml:"referrer,omitempty"`
	ViolatedDirective  string `json:"violated-directive" yaml:"violated-directive,omitempty"`
	EffectiveDirective string `json:"effective-directive" yaml:"effective-directive,omitempty"`
	OriginalPolicy     string `json:"original-policy" yaml:"original-policy,omitempty"`
	Disposition        string `json:"disposition" yaml:"disposition,omitempty"`
	BlockedURI         string `json:"blocked-uri" yaml:"blocked-uri,omitempty"`
	StatusCode         int    `json:"status-code" yaml:"status-code,omitempty"`
	ScriptSample       string `json:"script-sample" yaml:"script-sample,omitempty"`
}

func (h Handler) ReceiveContentSecurityPolicyReport(w http.ResponseWriter, r *http.Request) {
	raw, _ := io.ReadAll(r.Body)
	var request ContentSecurityPolicyRequest
	_ = json.Unmarshal(raw, &request)

	report := request.CSPReport
	reportJSON, _ := json.MarshalIndent(report, "", "  ")
	log.Println(string(reportJSON))

	if h.seed.Bucket.Name != "" {
		fmt.Println("--- bucket start ---")
		reportYAML, _ := yaml.Marshal(report)
		now := time.Now()

		key := fmt.Sprintf("reports/csp/%d/%d/%d/%d.yaml",
			now.Year(), now.Month(), now.Day(), now.Unix())

		_, err := s3.NewFromConfig(h.awsConfig).PutObject(r.Context(), &s3.PutObjectInput{
			Bucket: jsii.String(h.seed.Bucket.Name),
			Key:    jsii.String(key),
			Body:   bytes.NewReader(reportYAML),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created s3://%s/%s\n", h.seed.Bucket.Name, key)
		fmt.Println("--- bucket stop ---")
	}

	w.WriteHeader(201)
}
