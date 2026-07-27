package main

import (
	"log"
	"testing"

	"github.com/renatopp/cli-serve/serve"
	"github.com/renatopp/go-x/httpx/fetch"
	"github.com/renatopp/go-x/testx"
)

func TestStaticModeCLI(t *testing.T) {

	t.Run("serves files from directory", func(t *testing.T) {
		server := serve.NewStaticServer(serve.StaticServerOptions{
			Address:     ":0",
			Directory:   "./files",
			Logger:      log.New(nil, "", 0),
			Prefix:      "//sample/",
			SpaFallback: false,
			EnableCors:  false,
		})

		baseURL, closeServer := serveOnRandomPort(t, server)
		defer closeServer()

		res := fetch.Get(baseURL + "/sample/index.html")
		testx.Equal(t, 200, res.StatusCode)
		testx.Equal(t, "index mock\n", res.Text())

		res2 := fetch.Get(baseURL + "/sample/sample.md")
		testx.Equal(t, 200, res2.StatusCode)
		testx.Equal(t, "sample md\n", res2.Text())

		res3 := fetch.Get(baseURL + "/sample/asdf")
		testx.Equal(t, 404, res3.StatusCode)
	})

	t.Run("supports spa fallback", func(t *testing.T) {
		server := serve.NewStaticServer(serve.StaticServerOptions{
			Address:     ":0",
			Directory:   "./files",
			Logger:      log.New(nil, "", 0),
			Prefix:      "//sample/",
			SpaFallback: true,
			EnableCors:  false,
		})

		baseURL, closeServer := serveOnRandomPort(t, server)
		defer closeServer()

		res := fetch.Get(baseURL + "/sample/asdf")
		testx.Equal(t, 200, res.StatusCode)
		testx.Equal(t, "index mock\n", res.Text())

		res2 := fetch.Get(baseURL + "/sample/sample.md")
		testx.Equal(t, 200, res2.StatusCode)
		testx.Equal(t, "sample md\n", res2.Text())
	})

	t.Run("supports optional cors headers", func(t *testing.T) {
		server := serve.NewStaticServer(serve.StaticServerOptions{
			Address:     ":0",
			Directory:   "./files",
			Logger:      log.New(nil, "", 0),
			Prefix:      "//sample/",
			SpaFallback: false,
			EnableCors:  true,
		})

		baseURL, closeServer := serveOnRandomPort(t, server)
		defer closeServer()

		res := fetch.Get(baseURL + "/sample/")
		testx.Equal(t, 200, res.StatusCode)
		testx.Equal(t, "index mock\n", res.Text())
		testx.Equal(t, "*", res.Header("Access-Control-Allow-Origin"))
	})
}
