//with the suffix _test, this file will be ignored in go build.
//it will be taken when we run go test
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"campoy@golang.org", "gopher campoy"},
		{"something", "dear something"},
	}

	for _, c := range cases {
		// we need a http request
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil, //body
		)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
			return
		}

		// we need a fake ResponseWriter
		rec := httptest.NewRecorder() // rec is a ResponseWriter
		handler(rec, req)

		//now we want to make sure the response we got is actually correct
		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}

}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// we need a http request
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/campoy@golang.org",
			nil, //body
		)
		if err != nil {
			b.Fatalf("Could not create request: %v", err)
			return
		}

		// we need a fake ResponseWriter
		rec := httptest.NewRecorder() // rec is a ResponseWriter
		handler(rec, req)

		//now we want to make sure the response we got is actually correct
		if rec.Code != http.StatusOK {
			b.Errorf("expected status 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "gopher campoy") {
			b.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}

}
