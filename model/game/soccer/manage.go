package soccer

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	soccer matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	soccer.Matches = matche.InitGraph(teams)
	soccer.Matches = matche.UpdateGraph(soccer.Matches)
}
