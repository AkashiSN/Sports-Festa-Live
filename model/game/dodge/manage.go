package dodge

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

func Manage(teams []string) {
	var dodge Manager
	dodge.Matches = matche.InitMatch(teams)
	dodge.Matches = matche.UpdateSecondLayer(teams, dodge.Matches)
	fmt.Print(dodge.Matches)
}
