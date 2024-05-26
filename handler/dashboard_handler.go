package handler

import (
	"log"
	"net/http"

	"github.com/moroz/webauthn-academy-go/templates/dashboard"
)

type dashboardHandler struct{}

func DashboardHandler() dashboardHandler {
	return dashboardHandler{}
}

func (h *dashboardHandler) Index(w http.ResponseWriter, r *http.Request) {
	err := dashboard.Home().Render(r.Context(), w)
	if err != nil {
		log.Print(err)
	}
}
