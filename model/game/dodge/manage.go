package dodge

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	dodge matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	dodge.Matches = matche.UpdateFirstLayer(teams)
	dodge.Matches = matche.UpdateSecondLayer(teams, dodge.Matches)
	fmt.Print(dodge.Matches)
}
