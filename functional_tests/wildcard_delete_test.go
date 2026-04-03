// nolint dupl
package functional_tests

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/voronelf/httpfake"
)

// TestWildcardDelete tests a fake server handling a DELETE request with path parameter and using wildcard in path
func TestWildcardDelete(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	// register a handler for our fake service
	fakeService.NewHandler().
		Delete("/users/*/").
		Reply(200)

	req, err := http.NewRequest("DELETE", fakeService.ResolveURL("/users/1/"), nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close() // nolint errcheck

	// Check the status code is what we expect
	if status := res.StatusCode; status != 200 {
		t.Errorf("request returned wrong status code: got %v want %v",
			status, 200)
	}

	// Check the response body is what we expect
	expected := ""
	body, _ := ioutil.ReadAll(res.Body)
	if bodyString := string(body); bodyString != expected {
		t.Errorf("request returned unexpected body: got %v want %v",
			bodyString, expected)
	}
}
