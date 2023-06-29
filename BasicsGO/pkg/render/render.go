package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error message : ", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in cache
	_, inMap := tc[t]
	if !inMap {
		//need to create a template
		log.Println("creating template and add it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we already have it
		log.Println("using cache template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	} else {
		tc[t] = tmpl
	}
	return nil
}
