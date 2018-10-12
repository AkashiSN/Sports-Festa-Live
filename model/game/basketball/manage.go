package basketball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	basketball matche.Manager
)

func Manage(teams []string) {
	basketball.Matches = matche.InitMatch(teams)
	basketball.Matches = matche.UpdateSecondLayer(teams, basketball.Matches)
	fmt.Print(basketball.Matches)
}
