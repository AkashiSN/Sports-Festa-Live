package soccer

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var soccer Soccer
	soccer.Matches = matche.InitMatch(teams)
	soccer.Matches = matche.UpdateSecondLayer(teams, soccer.Matches)
	fmt.Print(soccer.Matches)
}
