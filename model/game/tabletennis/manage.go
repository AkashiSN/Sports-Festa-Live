package tabletennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var tabletennis Manager
	tabletennis.Matches = matche.InitMatch(teams)
	tabletennis.Matches = matche.UpdateSecondLayer(teams, tabletennis.Matches)
	fmt.Print(tabletennis.Matches)
}
