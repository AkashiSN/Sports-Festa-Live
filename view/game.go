package view

import (
	"net/http"
	"strconv"

	"github.com/AkashiSN/Sports-Festa-Live/model/game/badminton"
	"github.com/labstack/echo"
)

type jsonGame struct {
	NowPlay  []int `json:"now_play"`
	NextPlay []int `json:"next_play"`
}

type jsonMatch struct {
	TeamA  string `json:"team_a"`
	TeamB  string `json:"team_b"`
	Winner string `json:"winner"`
}

// GetBadminton バドミントンの試合の状況を取得する
func GetBadminton(e echo.Context) error {
	parm := e.Param("matchNum")
	if parm == "" {
		InCourtMatchNums, QueueMatchNums := badminton.GetMatchNums()
		info := jsonGame{
			NowPlay:  InCourtMatchNums,
			NextPlay: QueueMatchNums,
		}
		return e.JSON(http.StatusOK, info)
	} else {
		matchNum, err := strconv.Atoi(parm)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"error": err})
		}
		match := badminton.GetMatch(matchNum)
		info := jsonMatch{
			TeamA:  match.TeamA.Name,
			TeamB:  match.TeamB.Name,
			Winner: match.Winner,
		}
		return e.JSON(http.StatusOK, info)
	}
}
