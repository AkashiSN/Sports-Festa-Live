package j_dodge

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var j_dodge J_dodge
	j_dodge.Matches = matche.InitMatch(teams)
	j_dodge.Matches = matche.UpdateSecondLayer(teams, j_dodge.Matches)
	fmt.Print(j_dodge.Matches)
}
