package whatever

import (
	"encoding/json"
	"net/http"

	"github.com/jmulhern/whatever/pkg/kind"
)

func (h Handler) GetThings(w http.ResponseWriter, _ *http.Request) {
	raw, _ := json.Marshal([]kind.Thing{
		{Name: "Hello"},
	})
	_, _ = w.Write(raw)
}
