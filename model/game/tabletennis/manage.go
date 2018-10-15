package tabletennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tabletennis matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string, times []matche.MatcheTime) {
	tabletennis.Matches = matche.InitGraph(teams)
	tabletennis.Matches = matche.UpdateGraph(tabletennis.Matches)
}
