package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

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
		// if the template doesn't exist create a new template
		app.infoLogger.Printf("Caching template %s",name)
		app.serverError(w,fmt.Errorf("the template %s does not exists",name))
		return
	}

	// Some more logs
	app.infoLogger.Printf("Rendering the template %v",ts)

	buf := new(bytes.Buffer)

	// check to see what is in the buffer

	// TODO change the data being passed to the `Execute` method below
	err := ts.Execute(buf,td)
	if err != nil {
		app.serverError(w,err)
		return
	}

	n, err := buf.WriteTo(w)
	if err != nil {
		app.serverError(w,err)
		return
	}

	app.infoLogger.Printf("Rendered %d bytes",n)
}

