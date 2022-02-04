package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Initialize a new response recorder
	rr := httptest.NewRecorder()

	// Initialize a dummy http request
	r,err := http.NewRequest("GET","/",nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr,r)

	//Inspect the response recorder
	rs := rr.Result()

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
