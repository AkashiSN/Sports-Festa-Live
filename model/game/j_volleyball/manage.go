package j_volleyball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var j_volleyball Manager
	j_volleyball.Matches = matche.InitMatch(teams)
	j_volleyball.Matches = matche.UpdateSecondLayer(teams, j_volleyball.Matches)
	fmt.Print(j_volleyball.Matches)
}
