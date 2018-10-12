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
	matches, layer := matche.UpdateFirstLayer(teams)
	badminton.Matches = matches
	badminton.Layer = append(badminton.Layer, layer)
	badminton.Matches = matche.UpdateSecondLayer(teams, badminton.Matches, badminton.Layer[0])
	fmt.Print(badminton.Matches)
}
