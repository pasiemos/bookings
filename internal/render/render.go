package render

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/justinas/nosurf"
	"github.com/pasiemos/bookings/pkg/config"

	//"html_template_go/pkg/handlers"
	"log"
	"net/http"
	"path/filepath"

	"github.com/pasiemos/bookings/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

//NewTemplates sets the config for the template 
func NewTemplates(a * config.AppConfig)  {
	app = a
}

func AddDefaultData (td * models.TemplateData, r *http.Request) *models.TemplateData  {
	td.CSRFToken = nosurf.Token(r)
	return td
	
}

//RenderTemplate renders templates using html/template
//td template data ,import as a third parameter. Is of type handlers.TemplateData
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	//create a bytes buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing tremplate to browser", err)
	}
}

//CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error){
//make a map
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.tmpl.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
	
		// create a new variable ts, stands for template set 
		ts,  err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl.html")
			if err != nil{
				return myCache, err
			}
		}
		//take the template set that we created and add it to the Cache, that we created before
		//name is the actual name of the template
		myCache[name] = ts
	}
	return myCache, nil
}