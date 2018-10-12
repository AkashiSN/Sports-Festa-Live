package badminton

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	badminton matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	badminton.Matches = matche.UpdateFirstLayer(teams)
	badminton.Matches = matche.UpdateSecondLayer(teams, badminton.Matches)
	fmt.Print(badminton.Matches)
}
