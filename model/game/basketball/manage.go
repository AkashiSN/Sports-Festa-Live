package basketball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	basketball match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	basketball.Matches = match.InitGraph(teams)
	basketball.Matches = match.UpdateGraph(basketball.Matches)
}
