package j_volleyball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	j_volleyball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	j_volleyball.Matches = matche.UpdateFirstLayer(teams)
	j_volleyball.Matches = matche.UpdateSecondLayer(teams, j_volleyball.Matches)
	fmt.Print(j_volleyball.Matches)
}
