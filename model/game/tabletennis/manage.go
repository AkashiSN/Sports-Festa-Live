package tabletennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tabletennis matche.Manager
)

func Manage(teams []string) {
	tabletennis.Matches = matche.InitMatch(teams)
	tabletennis.Matches = matche.UpdateSecondLayer(teams, tabletennis.Matches)
	fmt.Print(tabletennis.Matches)
}
