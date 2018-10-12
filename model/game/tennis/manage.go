package tennis

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	tennis matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	matches, layer := matche.UpdateFirstLayer(teams)
	tennis.Matches = matches
	tennis.Layer = append(tennis.Layer, layer)
	tennis.Matches = matche.UpdateSecondLayer(teams, tennis.Matches, tennis.Layer[0])
	tennis.CurrentLayer = 0
	fmt.Println("tennis")
	for _, matche := range tennis.Matches {
		if matche.MatcheNum == 0 && matche.TeamA == "" {
			break
		}
		fmt.Printf("No.%d \"%s\" VS \"%s\"\n", matche.MatcheNum, matche.TeamA, matche.TeamB)
	}
}
