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
	softball.CurrentLayer = 0
	fmt.Println("softball")
	for _, matche := range softball.Matches {
		if matche.MatcheNum == 0 && matche.TeamA == "" {
			break
		}
		fmt.Printf("No.%d \"%s\" VS \"%s\"\n", matche.MatcheNum, matche.TeamA, matche.TeamB)
	}
}
