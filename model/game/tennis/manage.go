package tennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	tennis match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	tennis.Matches = match.InitGraph(teams)
	tennis.Matches = match.UpdateGraph(tennis.Matches)
}
