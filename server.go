package main

import (
	"net"
	"os"

	"github.com/AkashiSN/Sports-Festa-Live/view"

	"github.com/labstack/echo"
)

// 各メソッドの定義
const (
	GET int = iota
	POST
	PUT
	DELETE
)

// Route エンドポイント管理
type Route struct {
	method  int
	path    string
	handler echo.HandlerFunc
}

// APIのルーティングの指定
var routeAPI = []Route{
	{GET, "/api/game", view.GetGame},
}

// startServer サーバーを起動する
func startServer(sock bool) {
	e := echo.New()
	e.HideBanner = true

	// unix sockでlistenするか
	if sock {
		os.Remove("/tmp/echo.sock")
		l, err := net.Listen("unix", "/tmp/echo.sock")
		if err != nil {
			e.Logger.Fatal(err)
		}
		e.Listener = l
	} else {
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			e.Logger.Fatal(l)
		}
		e.Listener = l
	}

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
