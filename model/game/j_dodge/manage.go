package j_dodge

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	j_dodge matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	j_dodge.Matches = matche.UpdateFirstLayer(teams)
	j_dodge.Matches = matche.UpdateSecondLayer(teams, j_dodge.Matches)
	fmt.Print(j_dodge.Matches)
}
