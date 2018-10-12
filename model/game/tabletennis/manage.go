package tabletennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tabletennis matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	matches, layer := matche.UpdateFirstLayer(teams)
	tabletennis.Matches = matches
	tabletennis.Layer = append(tabletennis.Layer, layer)
	tabletennis.Matches = matche.UpdateSecondLayer(teams, tabletennis.Matches, tabletennis.Layer[0])
	fmt.Print(tabletennis.Matches)
}
