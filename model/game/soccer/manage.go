package soccer

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	soccer matche.Manager
)

func Manage(teams []string) {
	soccer.Matches = matche.InitMatch(teams)
	soccer.Matches = matche.UpdateSecondLayer(teams, soccer.Matches)
	fmt.Print(soccer.Matches)
}
