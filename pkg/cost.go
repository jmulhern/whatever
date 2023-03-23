package whatever

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"

	"github.com/jmulhern/whatever/pkg/kind"
)

func (h Handler) GetCosts(w http.ResponseWriter, _ *http.Request) {
	_ = costexplorer.NewFromConfig(h.awsConfig)
	raw, _ := json.Marshal([]kind.Cost{
		{Name: "ECR", Amount: "$100.00"},
		{Name: "ECS", Amount: "$10.00"},
		{Name: "Total", Amount: "$110.00"},
	})
	_, _ = w.Write(raw)
}
