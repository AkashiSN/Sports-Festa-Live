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
	matches, layer := matche.UpdateFirstLayer(teams)
	basketball.Matches = matches
	basketball.Layer = append(basketball.Layer, layer)
	basketball.Matches = matche.UpdateSecondLayer(teams, basketball.Matches, basketball.Layer[0])
	fmt.Print(basketball.Matches)
}
