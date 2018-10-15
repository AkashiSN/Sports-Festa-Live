package softball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	softball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string, times []matche.MatcheTime) {
	softball.Matches = matche.InitGraph(teams)
	softball.Matches = matche.UpdateGraph(softball.Matches)
}
