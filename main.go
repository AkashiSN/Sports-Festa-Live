package main

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/game"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

func main() {
	team := team.LoadTeam()
	game.GameStart(team)
}
