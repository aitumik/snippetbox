package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("expected %q; got%q", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("expteced body to contain %v", "OK")
	}
}

func TestShowSnippet(t *testing.T) {
	// Create an instance of the application using `netTestApplication() function`
	app := newTestApplication(t)

	// Create a server
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippets/1", http.StatusOK, []byte("I never meant for any of this to happen")},
		{"Non existent ID", "/snippets/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippets/-10", http.StatusNotFound, nil},
		{"Empty ID", "/snippets/", http.StatusNotFound, nil},
		{"Trailing Slash", "/snippets/1/", http.StatusNotFound, nil},
		{"String ID", "/snippets/foo", http.StatusNotFound, nil},
		{"Decimal ID", "/snippets/1.3", http.StatusNotFound, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			if code != tt.wantCode {
				t.Errorf("expected %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("expected body to contain %q", tt.wantBody)
			}
		})
	}
}

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")

	//TODO fix extractCSRFToken
	csrfToken := extractCSRFToken(t, body)
	t.Log(csrfToken)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid submission", "Aitumik", "aitumik@protonmail.com", "validPa$$word", csrfToken, http.StatusCreated, []byte("somethign")},
		{"Invalid submission", "Saitama", "saitama@protonmail.com", "validPa$$word", csrfToken, http.StatusCreated, []byte("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.postForm(t, "user/signup", form)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %s to contain %q", body, tt.wantBody)
			}
		})
	}
}
