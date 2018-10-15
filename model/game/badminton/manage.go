package badminton

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	badminton matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string, times []matche.MatcheTime) {
	badminton.Matches = matche.InitGraph(teams)
	badminton.Matches = matche.UpdateGraph(badminton.Matches)
}
