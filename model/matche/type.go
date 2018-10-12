package matche

// 各試合の構造体
type Matche struct {
	MatcheNum int
	TeamA     string
	TeamB     string
	Winner    string
}

// 各競技ごとの管理をする構造体
type Manager struct {
	Matches []Matche
	InCourt struct {
		CourtA string
		CourtB string
	}
	Queue struct {
		CourtA string
		CourtB string
	}
	Layer []int
}
