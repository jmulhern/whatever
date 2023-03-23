package whatever

import (
	"log"
	"net/http"
)

func (h Handler) GetIndex(w http.ResponseWriter, _ *http.Request) {
	err := h.index.Execute(w, map[string]any{
		"name": h.seed.Name,
		"site": h.seed.Site,
	})
	if err != nil {
		log.Fatal(err)
	}
}
