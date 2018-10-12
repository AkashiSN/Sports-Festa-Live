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
	matches, layer := matche.UpdateFirstLayer(teams)
	j_dodge.Matches = matches
	j_dodge.Layer = append(j_dodge.Layer, layer)
	j_dodge.Matches = matche.UpdateSecondLayer(teams, j_dodge.Matches, j_dodge.Layer[0])
	fmt.Print(j_dodge.Matches)
}
