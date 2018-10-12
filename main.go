package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AkashiSN/Sports-Festa-Live/model/game"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

func main() {
	s := flag.Bool("s", false, "Listen unix socket: unix:///tmp/echo.sock")
	flag.Parse()
	if *s {
		go startServer(true)
	} else {
		go startServer(false)
	}

	team := team.LoadTeam()
	game.GameStart(team)

	quit := make(chan int)
	go catchSignal(quit)

	code := <-quit
	os.Exit(code)
}

func catchSignal(c chan int) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		s := <-signalChan
		switch s {
		case syscall.SIGINT:
			c <- 0
		default:
			log.Printf("catch signal:%d", s)
			c <- 1
		}
	}
}
