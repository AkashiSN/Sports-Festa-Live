package jdodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	jdodge match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.MatchTime) {
	jdodge.Matches = match.InitGraph(teams)
	jdodge.Matches = match.UpdateGraph(jdodge.Matches)
}
