package team

// Team 各試合のチームの構造体
type Team struct {
	Dodge       []string `json:"dodge"`
	Basketball  []string `json:"basketball"`
	Badminton   []string `json:"badminton"`
	Tabletennis []string `json:"tabletennis"`
	Tennis      []string `json:"tennis"`
	Softball    []string `json:"softball"`
	Soccer      []string `json:"soccer"`
	Jvolleyball []string `json:"J_volleyball"`
	Jdodge      []string `json:"J_dodge"`
}
