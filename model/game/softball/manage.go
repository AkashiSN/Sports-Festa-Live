package softball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	softball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	softball.Matches = matche.UpdateFirstLayer(teams)
	softball.Matches = matche.UpdateSecondLayer(teams, softball.Matches)
	fmt.Print(softball.Matches)
}
