package forms

//import (
//	"crypto/ed25519/internal/edwards25519/field"

//	"golang.org/x/text/message"
//)

type errors map[string] []string

//Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}


//Get returns the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}