package whatever

import "net/http"

func (h Handler) GetHealth(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("OK"))
}
