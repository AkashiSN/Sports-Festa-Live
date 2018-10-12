package main

import (
	"net"
	"os"

	"github.com/AkashiSN/Sports-Festa-Live/view"

	"github.com/labstack/echo"
)

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
	PUT
	DELETE
)

type Route struct {
	method  HTTPMethod
	path    string
	handler echo.HandlerFunc
}

var routeAPI = []Route{
	{GET, "/api/game", view.GetGame},
}

func startServer(host string) {
	e := echo.New()

	os.Remove("/tmp/echo.sock")
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Listener = l

	for _, route := range routeAPI {
		switch route.method {
		case GET:
			e.GET(route.path, route.handler)
		case POST:
			e.POST(route.path, route.handler)
		case PUT:
			e.PUT(route.path, route.handler)
		case DELETE:
			e.DELETE(route.path, route.handler)
		}
	}

	e.Static("/", "static")
	e.Logger.Fatal(e.Start(""))
}
