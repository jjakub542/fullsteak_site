package app

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

type templateRegistry struct {
	templates map[string]*template.Template
}

func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)

		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func Renderer() *templateRegistry {
	templates := make(map[string]*template.Template)

	templates["home.html"] = template.Must(template.ParseFiles("web/templates/home.html", "web/templates/base.html"))
	templates["portfolio.html"] = template.Must(template.ParseFiles("web/templates/portfolio.html", "web/templates/base.html"))
	templates["blog.html"] = template.Must(template.ParseFiles("web/templates/blog.html", "web/templates/base.html"))
	templates["article.html"] = template.Must(template.ParseFiles("web/templates/article.html", "web/templates/base.html"))

	templates["admin/login.html"] = template.Must(template.ParseFiles("web/templates/admin/login.html"))
	templates["admin/logout.html"] = template.Must(template.ParseFiles("web/templates/admin/logout.html"))
	templates["admin/statistics.html"] = template.Must(template.ParseFiles("web/templates/admin/statistics.html", "web/templates/admin/base.html"))
	templates["admin/home.html"] = template.Must(template.ParseFiles("web/templates/admin/home.html", "web/templates/admin/base.html"))
	templates["admin/article.html"] = template.Must(template.ParseFiles("web/templates/admin/article.html", "web/templates/admin/base.html"))

	return &templateRegistry{
		templates: templates,
	}
}
