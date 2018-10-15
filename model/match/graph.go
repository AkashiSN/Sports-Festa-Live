package match

import (
	"math"
)

// InitGraph 正規グラフを作成する
func InitGraph(teams []string) [][]Match {
	// 正規グラフを作成する
	LayerCount := int(math.Ceil(math.Log2(float64(len(teams)))))
	graph := make([][]Match, LayerCount)
	for i := 0; i < LayerCount; i++ {
		graph[i] = make([]Match, int(math.Pow(2, float64(i))))
	}

	// 試合番号を振る
	matchNum := 0
	for i := LayerCount - 1; i >= 0; i-- {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			graph[i][j].MatchNum = matchNum
			matchNum++
		}
	}

	// 子の座標の計算を行う
	for i := 0; i < LayerCount; i++ {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			if i != LayerCount-1 {
				graph[i][j].TeamA.ChildX = 2*(j+1) - 2
				graph[i][j].TeamA.ChildY = i + 1
				graph[i][j].TeamB.ChildX = 2*(j+1) - 1
				graph[i][j].TeamB.ChildY = i + 1
			}
		}
	}

	// チームを割り当てる（シードかどうかも）
	for j := 0; j < len(graph[LayerCount-1])-1; j++ {
		if teams[j] == "" {
			graph[LayerCount-1][j].IsSeed = true
		}
		graph[LayerCount-1][j].Winner = teams[j]
	}
	return graph
}

// UpdateGraph グラフの再計算を行う
func UpdateGraph(graph [][]Match) [][]Match {
	for i := 0; i < len(graph)-1; i++ {
		for j := 0; j < len(graph[i]); j++ {
			xa := graph[i][j].TeamA.ChildX
			ya := graph[i][j].TeamA.ChildY
			xb := graph[i][j].TeamB.ChildX
			yb := graph[i][j].TeamB.ChildY
			if graph[yb][xb].IsSeed {
				graph[i][j].Winner = graph[ya][xa].Winner
			} else {
				graph[i][j].TeamA.Name = graph[ya][xa].Winner
				graph[i][j].TeamB.Name = graph[yb][xb].Winner
			}
		}
	}
	return graph
}
