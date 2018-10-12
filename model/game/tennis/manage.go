package tennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tennis matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	tennis.Matches = matche.UpdateFirstLayer(teams)
	tennis.Matches = matche.UpdateSecondLayer(teams, tennis.Matches)
	fmt.Print(tennis.Matches)
}
