package soccer

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	soccer matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	soccer.Matches = matche.UpdateFirstLayer(teams)
	soccer.Matches = matche.UpdateSecondLayer(teams, soccer.Matches)
	fmt.Print(soccer.Matches)
}
