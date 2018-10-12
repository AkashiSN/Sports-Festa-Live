package tabletennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tabletennis matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	tabletennis.Matches = matche.UpdateFirstLayer(teams)
	tabletennis.Matches = matche.UpdateSecondLayer(teams, tabletennis.Matches)
	fmt.Print(tabletennis.Matches)
}
