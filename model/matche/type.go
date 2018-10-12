package matche

type Matche struct {
	MatcheNum int
	TeamA     string
	TeamB     string
	Winner    string
}

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
}
