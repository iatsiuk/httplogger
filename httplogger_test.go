package httplogger

import (
	"net/http"
	"testing"

	"example.com/httplogger/adapters"
)

func TestClient(t *testing.T) {
	NewHTTPLogger(adapters.NewLogrusAdapter())

	client := &http.Client{}
	resp, err := client.Get("https://ya.ru")
	if err != nil {
		t.Errorf("TestClient: %s", err.Error())
	}
	defer resp.Body.Close()
}
