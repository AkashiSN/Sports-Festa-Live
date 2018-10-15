package softball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	softball match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	softball.Matchs = match.InitGraph(teams)
	softball.Matchs = match.UpdateGraph(softball.Matchs)
}
