package tennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	tennis match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	tennis.Matchs = match.InitGraph(teams)
	tennis.Matchs = match.UpdateGraph(tennis.Matchs)
}
