package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

var TemplateError = errors.New("template : template not found")

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s",err.Error(),debug.Stack())
	app.errorLogger.Output(2,trace)
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter,statusCode int) {
	http.Error(w,http.StatusText(statusCode),statusCode)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w,http.StatusNotFound)
}

// render `render the templates with the data provided`
func(app *application) render(w http.ResponseWriter,r *http.Request,name string,td *TemplateData)  {
	ts,ok := app.templateCache[name]
	if !ok  {
		app.serverError(w,fmt.Errorf("the template %s does not exists",name))
		return
	}

	buf := new(bytes.Buffer)

	// check to see what is in the buffer

	// TODO change the data being passed to the `Execute` method below
	err := ts.Execute(buf,app.addDefault(td,r))
	if err != nil {
		app.serverError(w,err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w,err)
		return
	}
}

func (app *application) addDefault(td *TemplateData,r *http.Request) *TemplateData {
	if td == nil {
		td = &TemplateData{}
	}

	//add this to check if the user is authenticated
	td.AuthenticatedUser = app.authenticatedUser(r)

	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r,"flash")
	return td
}

func (app *application) authenticatedUser(r *http.Request) int {
	return app.session.GetInt(r,"userID")
}

