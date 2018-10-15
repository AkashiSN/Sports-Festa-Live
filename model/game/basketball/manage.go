package basketball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	basketball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string, times []matche.MatcheTime) {
	basketball.Matches = matche.InitGraph(teams)
	basketball.Matches = matche.UpdateGraph(basketball.Matches)
}
