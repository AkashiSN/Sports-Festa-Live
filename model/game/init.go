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
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

// InitMatche 試合の準備を行う
func InitMatche(team team.Team, time matche.Time) {
	badminton.InitMatche(team.Badminton, time.Badminton)
	basketball.InitMatche(team.Basketball, time.Basketball)
	dodge.InitMatche(team.Dodge, time.Dodge)
	jdodge.InitMatche(team.Jdodge, time.Jdodge)
	jvolleyball.InitMatche(team.Jvolleyball, time.Jvolleyball)
	soccer.InitMatche(team.Soccer, time.Soccer)
	softball.InitMatche(team.Softball, time.Softball)
	tabletennis.InitMatche(team.Tabletennis, time.Tabletennis)
	tennis.InitMatche(team.Tennis, time.Tennis)
}
