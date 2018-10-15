package jvolleyball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	jvolleyball match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	jvolleyball.Matchs = match.InitGraph(teams)
	jvolleyball.Matchs = match.UpdateGraph(jvolleyball.Matchs)
}
