package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/wandirahman13/bookings/pkg/config"
	"github.com/wandirahman13/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	log.Println("[RENDER] get app value from main to app in render")
	app = a
}

// AddDefaultData is to pass some data to all pages
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {
	var tmpl_cache map[string]*template.Template

	// check if cache config is enabled
	log.Println("[RENDER] check if cache config is enabled")
	if app.UseCache {
		// create a template cache
		log.Println("[RENDER] cache is enabled, get all template cache from app.TemplateCache in render")
		tmpl_cache = app.TemplateCache
	} else {
		log.Println("[RENDER] cache is disabled, generate new template cache to app.TemplateCache in render")
		tmpl_cache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	log.Printf("[RENDER] get %s template from cache", t)
	tmpl, ok := tmpl_cache[t]
	if !ok {
		log.Fatal("[RENDER] could not get template from cache")
	}

	buff := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = tmpl.Execute(buff, td)

	// render the template
	log.Printf("[RENDER] render %s template", t)
	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	log.Println("[RENDER] CreateTemplateCache func is triggered")
	// create an empty map string that holds pointers to template.Template
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		// return error if no matching file found
		return myCache, err
	}

	// range through all files endig with *.page.tmpl
	for _, page := range pages {
		// get file name only for *.page.tmpl
		page_name := filepath.Base(page)
		// parse file *.page.tmpl to a template using page_name as a template name
		tmpl_page, err := template.New(page_name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// looking for any layouts
		layout, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// if layout files found, pass it to tmpl_page since the page is need layout to works
		if len(layout) > 0 {
			tmpl_page, err = tmpl_page.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// store tmpl_page to myCache with page_name as an index of the map
		myCache[page_name] = tmpl_page
	}

	// return all templates as a map and give no error after done looping through all files.
	return myCache, nil
}
