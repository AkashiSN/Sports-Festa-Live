package basketball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var basketball Manager
	basketball.Matches = matche.InitMatch(teams)
	basketball.Matches = matche.UpdateSecondLayer(teams, basketball.Matches)
	fmt.Print(basketball.Matches)
}
