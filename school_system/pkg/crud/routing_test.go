package crud_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"school/pkg/crud"
	"testing"
)

func TestRouting(t *testing.T) {
	server := httptest.NewServer(crud.Routing())
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/studentById?stuId=123", server.URL))

	if err != nil {
		t.Fatalf("could ot send GET request: %v", err)
	}

	//// Now we need to read response body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}
	response := string(bytes.TrimSuffix(b, []byte("\n")))
	if response == "404 page not found" {
		t.Fatalf("Page not found")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, res.StatusCode)
	}
}
