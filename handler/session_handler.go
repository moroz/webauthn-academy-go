package handler

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/handler/templates"
	"github.com/moroz/webauthn-academy-go/service"
)

type sessionHandler struct {
	us service.UserService
}

func SessionHandler(db *sqlx.DB) sessionHandler {
	return sessionHandler{service.NewUserService(db)}
}

type sessionsNewAssigns struct {
	RequestContext
}

func (h *sessionHandler) New(w http.ResponseWriter, r *http.Request) {
	err := templates.Sessions.New.Execute(w, sessionsNewAssigns{
		RequestContext: RequestContext{
			Title: "Sign in",
		},
	})
	if err != nil {
		log.Print(err)
	}
}
