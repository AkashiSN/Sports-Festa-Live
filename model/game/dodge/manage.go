package dodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	dodge match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	dodge.Matches = match.InitGraph(teams)
	dodge.Matches = match.UpdateGraph(dodge.Matches)
}
