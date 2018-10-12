package basketball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	basketball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	basketball.Matches = matche.UpdateFirstLayer(teams)
	basketball.Matches = matche.UpdateSecondLayer(teams, basketball.Matches)
	fmt.Print(basketball.Matches)
}
