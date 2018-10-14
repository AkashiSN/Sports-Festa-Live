package matche

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
