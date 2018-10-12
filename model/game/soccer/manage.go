package soccer

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	soccer matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	matches, layer := matche.UpdateFirstLayer(teams)
	soccer.Matches = matches
	soccer.Layer = append(soccer.Layer, layer)
	soccer.Matches = matche.UpdateSecondLayer(teams, soccer.Matches, soccer.Layer[0])
	soccer.CurrentLayer = 0
	fmt.Println("soccer")
	for _, matche := range soccer.Matches {
		if matche.MatcheNum == 0 && matche.TeamA == "" {
			break
		}
		fmt.Printf("No.%d \"%s\" VS \"%s\"\n", matche.MatcheNum, matche.TeamA, matche.TeamB)
	}
}
