package soccer

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	soccer match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	soccer.Matchs = match.InitGraph(teams)
	soccer.Matchs = match.UpdateGraph(soccer.Matchs)
}
