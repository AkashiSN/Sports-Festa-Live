package game

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/game/badminton"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/basketball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/jdodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/jvolleyball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/soccer"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/softball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tabletennis"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tennis"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

// InitMatche 試合の準備を行う
func InitMatche(team team.Team) {
	badminton.InitMatche(team.Badminton)
	basketball.InitMatche(team.Basketball)
	dodge.InitMatche(team.Dodge)
	jdodge.InitMatche(team.Jdodge)
	jvolleyball.InitMatche(team.Jvolleyball)
	soccer.InitMatche(team.Soccer)
	softball.InitMatche(team.Softball)
	tabletennis.InitMatche(team.Tabletennis)
	tennis.InitMatche(team.Tennis)
}
