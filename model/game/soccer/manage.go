package soccer

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	soccer match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	soccer.Matches = match.InitGraph(teams)
	soccer.Matches = match.UpdateGraph(soccer.Matches)
}
