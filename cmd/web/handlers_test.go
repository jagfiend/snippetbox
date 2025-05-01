package main

import (
	"net/http"
	"testing"

	"github.com/jagfiend/snippetbox/internal/assert"
)

// example using our testing utils
func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())

	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

// simple unit test example using httptest.NewRecorder()
// func TestPing(t *testing.T) {
// 	rr := httptest.NewRecorder()

// 	r, err := http.NewRequest(http.MethodGet, "/", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ping(rr, r)

// 	rs := rr.Result()

// 	assert.Equal(t, rs.StatusCode, http.StatusOK)

// 	defer rs.Body.Close()

// 	body, err := io.ReadAll(rs.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	body = bytes.TrimSpace(body)

// 	assert.Equal(t, string(body), "OK")
// }
