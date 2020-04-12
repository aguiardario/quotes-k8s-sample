package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetConnectionString(t *testing.T) {

	urlExpected := "mongodb://localhost:27017/quotes"

	url := getConnectionString()
	if url != urlExpected {
		t.Error("Expected", urlExpected, "Got", url)
	}

	os.Setenv("HOST", "goole.com")
	os.Setenv("PORT", "27018")

	urlExpected = "mongodb://goole.com:27018/quotes"
	url = getConnectionString()
	if url != urlExpected {
		t.Error("Expected", urlExpected, "Got", url)
	}

}

func TestHome(t *testing.T) {

	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "27017")
	os.Setenv("USER_DB", "sapo")
	os.Setenv("PWD_DB", "123456")

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	nonExpected := `{}`
	if rr.Body.String() == nonExpected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), nonExpected)
	}
}
