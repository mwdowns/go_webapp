package render

import (
	"bytes"
	"fmt"
	"go_webapp/pkg/config"
	"go_webapp/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the new package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData sets the default data for pages
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// renderTemplate to renter templates for routes
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("no template of that name")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// createTemplateCache does what it says on the tin
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	handleError(myCache, err)

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		handleError(myCache, err)

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		handleError(myCache, err)

		if len(matches) > 0 {
			_, err := ts.ParseGlob("./templates/*.layout.tmpl")
			handleError(myCache, err)
		}
		myCache[name] = ts
	}
	return myCache, nil
}

func handleError(c map[string]*template.Template, e error) (map[string]*template.Template, error) {
	if e != nil {
		fmt.Println("Error:", e)
		return c, e
	}
	return c, nil
}
