package main

import (
	"fmt"
	"github.com/pasiemos/bookings/internal/config"
	"github.com/pasiemos/bookings/internal/handlers"
	"github.com/pasiemos/bookings/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
//var session is going to be a pointer to scs.SessionManager
var session *scs.SessionManager

//main is the main application function
func main() {
	

	//change this to true when in production
	app.InProduction = false


	//create a variable called session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//my template cache tc and an error is equal to 
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		
	}
	app.TemplateCache = tc
	app.UseCache = false

	//repo is the repository variable
	//handlers.NewRepo function. I'm gona pass it a reference to my app
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

			//(it wants a pointer) we are going to reference to that ...&
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))


	//srv serving
	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	//start the actual server
	err = srv.ListenAndServe()
	log.Fatal(err)
}
