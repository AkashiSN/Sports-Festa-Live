package badminton

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	badminton match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	badminton.Matches = match.InitGraph(teams)
	badminton.Matches = match.UpdateGraph(badminton.Matches)
}
