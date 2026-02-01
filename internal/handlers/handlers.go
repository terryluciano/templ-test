package handlers

import (
	"net/http"

	"github.com/terryluciano/templ-test/internal/views/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name := "John Doe"

	ip := r.RemoteAddr

	pages.Home(name, ip).Render(r.Context(), w)
}
