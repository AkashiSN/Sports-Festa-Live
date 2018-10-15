package badminton

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	badminton match.Manager
	graph     [][]match.Match
	numsIdx   []match.NumIdx
)

// GetInfo Managerを返す
func GetInfo() match.Manager {
	return badminton
}

// GetMatchNums 試合番号の配列を返す
func GetMatchNums() ([]int, []int) {
	return badminton.InCourt.MatchNums, badminton.Queue.MatchNums
}

// GetMatch 試合を返す
func GetMatch(matchNum int) match.Match {
	x := numsIdx[matchNum-1].X
	y := numsIdx[matchNum-1].Y
	return badminton.Matches[y][x]
}

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	graph, numsIdx = match.InitGraph(teams)
	graph = match.UpdateGraph(graph)
	graph = match.RegisterMatchTime(graph, times)
	match.DebugDump(graph)
	badminton.Matches = graph

	// キューに追加する
	badminton.InCourt.CourtA = GetMatch(1)
	badminton.InCourt.CourtB = GetMatch(2)
	badminton.InCourt.MatchNums = []int{GetMatch(1).MatchNum, GetMatch(2).MatchNum}
	badminton.Queue.CourtA = GetMatch(3)
	badminton.Queue.CourtB = GetMatch(4)
	badminton.Queue.MatchNums = []int{GetMatch(3).MatchNum, GetMatch(4).MatchNum}
}
