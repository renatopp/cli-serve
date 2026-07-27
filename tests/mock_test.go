package main

import (
	"io"
	"log"
	"testing"

	"github.com/renatopp/go-serve/serve"
	"github.com/renatopp/go-x/httpx/fetch"
	"github.com/renatopp/go-x/testx"
)

func TestMockModeCLI(t *testing.T) {
	t.Run("returns configured status", func(t *testing.T) {
		server := serve.NewMockServer(serve.MockServerOptions{
			Address: ":0",
			Logger:  log.New(io.Discard, "", 0),
			Status:  201,
		})

		baseURL, closeServer := serveOnRandomPort(t, server)
		defer closeServer()

		res := fetch.Get(baseURL + "/")
		testx.Equal(t, 201, res.StatusCode)

	})

	t.Run("returns configured body and headers", func(t *testing.T) {
		t.Skip("TODO: implement body/header assertions")
	})

	t.Run("applies response delay", func(t *testing.T) {
		t.Skip("TODO: implement delay assertions")
	})
}
