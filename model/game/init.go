package game

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/game/badminton"
	/*"github.com/AkashiSN/Sports-Festa-Live/model/game/basketball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/j_dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/j_volleyball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/soccer"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/softball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tabletennis"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tennis"*/
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

func InitMatche(team team.Team) {
	badminton.InitMatche(team.Badminton)
	/*	basketball.InitMatche(team.Basketball)
		dodge.InitMatche(team.Dodge)
		j_dodge.InitMatche(team.J_dodge)
		j_volleyball.InitMatche(team.J_volleyball)
		soccer.InitMatche(team.Soccer)
		softball.InitMatche(team.Softball)
		tabletennis.InitMatche(team.Tabletennis)
		tennis.InitMatche(team.Tennis)*/
}
