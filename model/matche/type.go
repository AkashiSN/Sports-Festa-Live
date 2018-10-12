package matche

type Matche struct {
	MatcheNum int
	TeamA     string
	TeamB     string
	Winner    string
}

type Manager struct {
	Matches []Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
