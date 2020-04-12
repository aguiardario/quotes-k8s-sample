package main

import (
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
