package badminton

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	badminton matche.Manager
)

func Manage(teams []string) {
	badminton.Matches = matche.InitMatch(teams)
	badminton.Matches = matche.UpdateSecondLayer(teams, badminton.Matches)
	fmt.Print(badminton.Matches)
}
