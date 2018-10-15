package game

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/game/badminton" /* 	"github.com/AkashiSN/Sports-Festa-Live/model/game/basketball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/dodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/jdodge"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/jvolleyball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/soccer"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/softball"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tabletennis"
	"github.com/AkashiSN/Sports-Festa-Live/model/game/tennis" */
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
	"github.com/AkashiSN/Sports-Festa-Live/model/team"
)

// InitMatch 試合の準備を行う
func InitMatch(team team.Teams, time match.Times) {
	badminton.InitMatch(team.Badminton, time.Badminton)
	/* 	basketball.InitMatch(team.Basketball, time.Basketball)
	   	dodge.InitMatch(team.Dodge, time.Dodge)
	   	jdodge.InitMatch(team.Jdodge, time.Jdodge)
	   	jvolleyball.InitMatch(team.Jvolleyball, time.Jvolleyball)
	   	soccer.InitMatch(team.Soccer, time.Soccer)
	   	softball.InitMatch(team.Softball, time.Softball)
	   	tabletennis.InitMatch(team.Tabletennis, time.Tabletennis)
	   	tennis.InitMatch(team.Tennis, time.Tennis) */
}
