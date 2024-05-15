package templates

import (
	"embed"
	"html/template"
)

//go:embed **/*.html.tmpl
var viewsFS embed.FS

func ParseTemplate(segments ...string) template.Template {
	newSegments := append([]string{"layout/root.html.tmpl"}, segments...)
	return *template.Must(template.ParseFS(viewsFS, newSegments...))
}

type UsersTemplates struct {
	New template.Template
}

var Users = UsersTemplates{
	New: ParseTemplate("users/new.html.tmpl"),
}

type SessionsTemplates struct {
	New template.Template
}

var Sessions = SessionsTemplates{
	New: ParseTemplate("sessions/new.html.tmpl"),
}
