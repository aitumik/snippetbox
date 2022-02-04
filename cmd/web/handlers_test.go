package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Create a buffer to act as writer
	//buf := new(bytes.Buffer)
	// Create an application instance
	app := &application{
		infoLogger: log.New(ioutil.Discard,"",0),
		errorLogger: log.New(ioutil.Discard,"",0),
	}

	// Create a new test tls server
	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	// Send a GET request to the test server
	rs,err:= ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("expected %q; got %q",http.StatusOK,rs.StatusCode)
	}

	defer rs.Body.Close()
	body,err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q","OK")
	}
}
