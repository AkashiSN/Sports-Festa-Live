package match

// Time 各競技の開始時刻と終了時刻の構造体
type Time struct {
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
}

// NumIdx 試合番号から[][]Matchへのインデクスの構造体
type NumIdx struct {
	X int
	Y int
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
		CourtA    Match
		CourtB    Match
		MatchNums []int
	}
	Queue struct {
		CourtA    Match
		CourtB    Match
		MatchNums []int
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
