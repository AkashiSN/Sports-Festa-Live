package dodge

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	dodge matche.Manager
)

func Manage(teams []string) {
	dodge.Matches = matche.InitMatch(teams)
	dodge.Matches = matche.UpdateSecondLayer(teams, dodge.Matches)
	fmt.Print(dodge.Matches)
}
