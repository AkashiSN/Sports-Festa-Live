package dodge

import (
	"fmt"

	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	dodge matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	matches, layer := matche.UpdateFirstLayer(teams)
	dodge.Matches = matches
	dodge.Layer = append(dodge.Layer, layer)
	dodge.Matches = matche.UpdateSecondLayer(teams, dodge.Matches, dodge.Layer[0])
	dodge.CurrentLayer = 0
	fmt.Println("dodge")
	for _, matche := range dodge.Matches {
		if matche.MatcheNum == 0 && matche.TeamA == "" {
			break
		}
		fmt.Printf("No.%d \"%s\" VS \"%s\"\n", matche.MatcheNum, matche.TeamA, matche.TeamB)
	}
}
