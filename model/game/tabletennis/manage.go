package tabletennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/match"
)

var (
	tabletennis match.Manager
)

// InitMatch 試合の準備を行う
func InitMatch(teams []string, times []match.Time) {
	graph := match.InitGraph(teams)
	graph = match.UpdateGraph(graph)
	graph = match.RegisterMatchTime(graph, times)
	tabletennis.Matches = graph
}
