package softball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	softball match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	softball.Matches = match.InitGraph(teams)
	softball.Matches = match.UpdateGraph(softball.Matches)
}
