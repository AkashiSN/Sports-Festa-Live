package j_volleyball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	j_volleyball matche.Manager
)

func Manage(teams []string) {
	j_volleyball.Matches = matche.InitMatch(teams)
	j_volleyball.Matches = matche.UpdateSecondLayer(teams, j_volleyball.Matches)
	fmt.Print(j_volleyball.Matches)
}
