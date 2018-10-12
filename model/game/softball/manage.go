package softball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var softball Manager
	softball.Matches = matche.InitMatch(teams)
	softball.Matches = matche.UpdateSecondLayer(teams, softball.Matches)
	fmt.Print(softball.Matches)
}
