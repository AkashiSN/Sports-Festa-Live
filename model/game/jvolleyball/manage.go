package jvolleyball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	jvolleyball match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	jvolleyball.Matches = match.InitGraph(teams)
	jvolleyball.Matches = match.UpdateGraph(jvolleyball.Matches)
}
