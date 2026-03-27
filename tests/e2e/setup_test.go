//go:build e2e

package e2e

import (
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	// create app app.CreateApp()

	// create server http.Server
	srv := &http.Server{}

	//listen and serve
	go srv.ListenAndServe()

	code := m.Run()

	srv.Close()

	os.Exit(code)
}
