package badminton

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var badminton Badminton
	badminton.Matches = matche.InitMatch(teams)
	badminton.Matches = matche.UpdateSecondLayer(teams, badminton.Matches)
	fmt.Print(badminton.Matches)
}
