package softball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	softball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	matches, layer := matche.UpdateFirstLayer(teams)
	softball.Matches = matches
	softball.Layer = append(softball.Layer, layer)
	softball.Matches = matche.UpdateSecondLayer(teams, softball.Matches, softball.Layer[0])
	fmt.Print(softball.Matches)
}
