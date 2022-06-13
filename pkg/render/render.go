package render

import (
	"bytes"
	"fmt"
	"github.com/mateusjunges/lets-go/pkg/config"
	"github.com/mateusjunges/lets-go/pkg/models"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates	sets the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = RegisterTemplateCache()
	}

	templatePage := templateCache[tmpl]

	templateBuffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = templatePage.Execute(templateBuffer, td)

	_, err := templateBuffer.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// RegisterTemplateCache creates a template cache as a map
func RegisterTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		templates, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(templates) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return cache, err
			}
		}

		cache[name] = templateSet
	}

	return cache, nil
}