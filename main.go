package main

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/game/dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

func main() {
	team := team.LoadTeam()
	dodge.Manage(team.Dodge)
}
