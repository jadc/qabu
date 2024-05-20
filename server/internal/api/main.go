package api

import (
	"net/http"
	"html/template"
)

// A struct that contains all templates defined in the views directory
type Templates struct { 
    templates *template.Template 
}

var templates *Templates

func GetTemplates() *Templates {
    if templates == nil {
        templates = &Templates{
            templates: template.Must(template.ParseGlob("views/*.html")),
        }
    }
    return templates
}

func (t *Templates) Render(w http.ResponseWriter, name string, data interface{}) error {
    err := t.templates.ExecuteTemplate(w, name, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    return err
}


