package models

import "github.com/pasiemos/bookings/internal/forms"

//TemplateData type holds data sent from handlers to templates
type TemplateData struct{
	//Data This type might be structs, might be things that hold all kinds all other data. It's going to be a map
	//A map with an index of string but the content of the map can be anything.
	//interface we are declaring it as a type, that's why the curely brckets after it
	
	StringMap 	map[string]string
	IntMap 		map[string]int
	FloatMap 	map[string]float32
	Data 		map[string]interface{}
	//CSRF Cross-site request forgery
	CSRFToken 	string
	Flash 		string
	Warning 	string
	Error 		string
	Form 		*forms.Form
}