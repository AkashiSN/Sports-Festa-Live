package tennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tennis matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	tennis.Matches = matche.InitGraph(teams)
	tennis.Matches = matche.UpdateGraph(tennis.Matches)
}
