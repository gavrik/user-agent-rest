package useragentrest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUaParser(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()
	resp, _ := http.Post(fmt.Sprintf("%s/ua", ts.URL), "application/x-www-form-urlencoded", strings.NewReader("localhost"))
	val, _ := resp.Header["Content-Type"]
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
