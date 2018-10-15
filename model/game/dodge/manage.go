package dodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	dodge match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	dodge.Matchs = match.InitGraph(teams)
	dodge.Matchs = match.UpdateGraph(dodge.Matchs)
}
