package jvolleyball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

var (
	jvolleyball matche.Manager
)

// InitMatche 試合の準備を行う
func InitMatche(teams []string) {
	jvolleyball.Matches = matche.InitGraph(teams)
	jvolleyball.Matches = matche.UpdateGraph(jvolleyball.Matches)
}
