package match

import (
	"time"
)

// customTime 独自のTime型
type customTime struct {
	time.Time
}

// Time 各競技の開始時刻と終了時刻の構造体
type Time struct {
	StartTime customTime `json:"start"`
	EndTime   customTime `json:"end"`
}

// Team チームの構造体
type Team struct {
	Name   string
	ChildX int
	ChildY int
}

// Match 各試合の構造体
type Match struct {
	MatchNum int
	Idx      int
	TeamA    Team
	TeamB    Team
	Winner   string
	IsSeed   bool
	Time     Time
}

// Manager 各競技ごとの管理をする構造体
type Manager struct {
	Matches      [][]Match
	Layer        []int
	CurrentLayer int

	InCourt struct {
		CourtA string
		CourtB string
	}
	Queue struct {
		CourtA string
		CourtB string
	}
}

// Times 各試合の競技時間の構造体
type Times struct {
	Dodge       []Time `json:"dodge"`
	Basketball  []Time `json:"basketball"`
	Badminton   []Time `json:"badminton"`
	Tabletennis []Time `json:"tabletennis"`
	Tennis      []Time `json:"tennis"`
	Softball    []Time `json:"softball"`
	Soccer      []Time `json:"soccer"`
	Jvolleyball []Time `json:"J_volleyball"`
	Jdodge      []Time `json:"J_dodge"`
}
