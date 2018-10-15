package matche

// MatcheTime 各競技の開始時刻と終了時刻の構造体
type MatcheTime struct {
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
}

// Team チームの構造体
type Team struct {
	Name   string
	ChildX int
	ChildY int
}

// Matche 各試合の構造体
type Matche struct {
	MatcheNum int
	TeamA     Team
	TeamB     Team
	Winner    string
	IsSeed    bool
	Time      MatcheTime
}

// Manager 各競技ごとの管理をする構造体
type Manager struct {
	Matches      [][]Matche
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
	Dodge       []MatcheTime `json:"dodge"`
	Basketball  []MatcheTime `json:"basketball"`
	Badminton   []MatcheTime `json:"badminton"`
	Tabletennis []MatcheTime `json:"tabletennis"`
	Tennis      []MatcheTime `json:"tennis"`
	Softball    []MatcheTime `json:"softball"`
	Soccer      []MatcheTime `json:"soccer"`
	Jvolleyball []MatcheTime `json:"J_volleyball"`
	Jdodge      []MatcheTime `json:"J_dodge"`
}
