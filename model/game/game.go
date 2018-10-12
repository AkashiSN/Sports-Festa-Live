package game

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/game/badminton"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/basketball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/j_dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/j_volleyball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/soccer"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/softball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tabletennis"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tennis"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

func GameStart(team team.Team) {
	badminton.Manage(team.Badminton)
	basketball.Manage(team.Basketball)
	dodge.Manage(team.Dodge)
	j_dodge.Manage(team.J_dodge)
	j_volleyball.Manage(team.J_volleyball)
	soccer.Manage(team.Soccer)
	softball.Manage(team.Softball)
	tabletennis.Manage(team.Tabletennis)
	tennis.Manage(team.Tennis)
}
