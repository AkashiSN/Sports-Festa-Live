package match

// MatchTime 各競技の開始時刻と終了時刻の構造体
type MatchTime struct {
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
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
	TeamA    Team
	TeamB    Team
	Winner   string
	IsSeed   bool
	Time     MatchTime
}

// Manager 各競技ごとの管理をする構造体
type Manager struct {
	Matchs       [][]Match
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

// Time 各試合の競技時間の構造体
type Time struct {
	Dodge       []MatchTime `json:"dodge"`
	Basketball  []MatchTime `json:"basketball"`
	Badminton   []MatchTime `json:"badminton"`
	Tabletennis []MatchTime `json:"tabletennis"`
	Tennis      []MatchTime `json:"tennis"`
	Softball    []MatchTime `json:"softball"`
	Soccer      []MatchTime `json:"soccer"`
	Jvolleyball []MatchTime `json:"J_volleyball"`
	Jdodge      []MatchTime `json:"J_dodge"`
}
