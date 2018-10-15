package tabletennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	tabletennis match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	tabletennis.Matchs = match.InitGraph(teams)
	tabletennis.Matchs = match.UpdateGraph(tabletennis.Matchs)
}
