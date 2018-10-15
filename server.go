package main

import (
	"net"
	"os"

	"github.com/AkashiSN/Sports-Festa-Live/view"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	//{GET, "/api/game", view.GetGame},
	{GET, "/api/game/badminton", view.GetBadminton},
	{GET, "/api/game/badminton/:matchNum", view.GetBadminton},
	//{PUT, "/api/game/badminton/:matchNum", view.SetBadminton},
}

// corsで許可するアドレス
var debugURL = []string{"http://localhost:8080"}

// startServer サーバーを起動する
func startServer(sock bool) {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     debugURL,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

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
	e.Start("")
}
