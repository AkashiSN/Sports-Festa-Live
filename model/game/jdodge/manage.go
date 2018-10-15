package jdodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	jdodge matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string, times []matche.MatcheTime) {
	jdodge.Matches = matche.InitGraph(teams)
	jdodge.Matches = matche.UpdateGraph(jdodge.Matches)
}
