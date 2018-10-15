package jdodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	jdodge match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	jdodge.Matchs = match.InitGraph(teams)
	jdodge.Matchs = match.UpdateGraph(jdodge.Matchs)
}
