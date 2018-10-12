package j_volleyball

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	j_volleyball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	matches, layer := matche.UpdateFirstLayer(teams)
	j_volleyball.Matches = matches
	j_volleyball.Layer = append(j_volleyball.Layer, layer)
	j_volleyball.Matches = matche.UpdateSecondLayer(teams, j_volleyball.Matches, j_volleyball.Layer[0])
	j_volleyball.CurrentLayer = 0
	fmt.Println("j_volleyball")
	for _, matche := range j_volleyball.Matches {
		if matche.MatcheNum == 0 && matche.TeamA == "" {
			break
		}
		fmt.Printf("No.%d \"%s\" VS \"%s\"\n", matche.MatcheNum, matche.TeamA, matche.TeamB)
	}
}
