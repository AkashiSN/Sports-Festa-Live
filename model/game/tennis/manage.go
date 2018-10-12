package tennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var tennis Tennis
	tennis.Matches = matche.InitMatch(teams)
	tennis.Matches = matche.UpdateSecondLayer(teams, tennis.Matches)
	fmt.Print(tennis.Matches)
}
