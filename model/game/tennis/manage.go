package tennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tennis matche.Manager
)

func Manage(teams []string) {
	tennis.Matches = matche.InitMatch(teams)
	tennis.Matches = matche.UpdateSecondLayer(teams, tennis.Matches)
	fmt.Print(tennis.Matches)
}
