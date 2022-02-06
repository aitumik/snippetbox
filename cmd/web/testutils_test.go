package main

import (
	"github.com/aitumik/snippetbox/pkg"
	"github.com/aitumik/snippetbox/pkg/models/mock"
	"github.com/golangcollege/sessions"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"
)

func newTestApplication(t *testing.T) *application {
	templateCache,err := NewTemplateCache("./../../ui/html")
	if err != nil {
		t.Fatal(err)
	}

	session := sessions.New([]byte("s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	var snippetModel = &mock.SnippetModel{}

	return &application{
		infoLogger: log.New(ioutil.Discard,"",0),
		errorLogger: log.New(ioutil.Discard,"",0),
		session: session,
		snippet: snippetModel,
		templateCache: templateCache,
		cfg: new(pkg.Config),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, handler http.Handler) *testServer{
	ts := httptest.NewTLSServer(handler)

	jar,err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	ts.Client().Jar = jar
	ts.Client().CheckRedirect = func(req *http.Request,via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T,urlPath string) (int,http.Header,[]byte) {
	rs,err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body,err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode,rs.Header,body
}

