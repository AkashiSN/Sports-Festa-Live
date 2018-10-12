package softball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	softball matche.Manager
)

func Manage(teams []string) {
	softball.Matches = matche.InitMatch(teams)
	softball.Matches = matche.UpdateSecondLayer(teams, softball.Matches)
	fmt.Print(softball.Matches)
}
