package dodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	dodge matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	dodge.Matches = matche.InitGraph(teams)
	dodge.Matches = matche.UpdateGraph(dodge.Matches)
}
