package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t,app.routes())
	defer ts.Close()

	code,_,body := ts.get(t,"/ping")

	if code != http.StatusOK {
		t.Errorf("expected %q; got%q",http.StatusOK,code)
	}

	if string(body) != "OK" {
		t.Errorf("expteced body to contain %v","OK")
	}

}
