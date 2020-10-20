package httplogger

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewHTTPLoggerDebugLevel(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	}))
	defer ts.Close()

	client := &http.Client{
		Transport: NewHTTPLogger(http.DefaultTransport, NewLogrusAdapter()),
	}
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Errorf("TestClient: %s", err.Error())
	}
	defer resp.Body.Close()
}

func TestNewHTTPLoggerTraceLevel(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	}))
	defer ts.Close()

	client := &http.Client{
		Transport: NewHTTPLogger(http.DefaultTransport, NewLogrusAdapter()),
	}
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Errorf("TestClient: %s", err.Error())
	}
	defer resp.Body.Close()
}
