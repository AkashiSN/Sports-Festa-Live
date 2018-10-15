package match

import (
	"math"
)

// InitGraph 正規グラフを作成する
func InitGraph(teams []string) [][]Match {
	// 正規グラフを作成する
	LayerCount := int(math.Ceil(math.Log2(float64(len(teams))))) + 1
	graph := make([][]Match, LayerCount)
	for i := 0; i < LayerCount; i++ {
		graph[i] = make([]Match, int(math.Pow(2, float64(i))))
	}

	// インデックスを振る
	idx := 0
	for i := LayerCount - 1; i >= 0; i-- {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			graph[i][j].Idx = idx
			idx++
		}
	}

	// 子の座標の計算を行う
	for i := 0; i < LayerCount-1; i++ {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			graph[i][j].TeamA.ChildX = 2*(j+1) - 2
			graph[i][j].TeamA.ChildY = i + 1
			graph[i][j].TeamB.ChildX = 2*(j+1) - 1
			graph[i][j].TeamB.ChildY = i + 1
		}
	}

	// チームを割り当てる（シードかどうかも）
	for j := 0; j < len(graph[LayerCount-1]); j++ {
		if teams[j] == "" {
			graph[LayerCount-1][j].IsSeed = true
		}
		graph[LayerCount-1][j].Winner = teams[j]
	}

	// 試合番号を振る
	matchCount := 1
	// トーナメント表の第1層について
	for i := 0; i < len(graph[LayerCount-1])-3; i += 4 {
		for j := i; j < i+3; j++ {
			if teams[j] == "" && teams[j+1] == "" {
				graph[LayerCount-2][j/2].MatchNum = matchCount
				graph[LayerCount-2][j/2+1].MatchNum = matchCount
				matchCount++
			}
			if j != i+1 && teams[j] != "" && teams[j+1] != "" {
				graph[LayerCount-2][j/2].MatchNum = matchCount
				matchCount++
			}
		}
	}
	// トーナメント表の第2層について
	count := 0
	for i := 0; i < len(graph[LayerCount-1])-3; i += 4 {
		if (teams[i+1] == "") != (teams[i+2] == "") {
			if teams[i+1] == "" {
				graph[LayerCount-3][count].MatchNum = matchCount
				matchCount++
			}
			if teams[i+2] == "" {
				graph[LayerCount-3][count].MatchNum = matchCount
				matchCount++
			}
		} else if !(teams[i+1] == "" && teams[i+2] == "") {
			graph[LayerCount-3][count].MatchNum = matchCount
			matchCount++
		}
		count++
	}
	// トーナメント表のそれ以上の層について
	for i := LayerCount - 4; i >= 0; i-- {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			graph[i][j].MatchNum = matchCount
			matchCount++
		}
	}

	// トーナメント表を更新する
	for i := 0; i < len(graph)-1; i++ {
		for j := 0; j < len(graph[i]); j++ {
			xa := graph[i][j].TeamA.ChildX
			ya := graph[i][j].TeamA.ChildY
			xb := graph[i][j].TeamB.ChildX
			yb := graph[i][j].TeamB.ChildY
			if graph[ya][xa].IsSeed {
				graph[i][j].Winner = graph[yb][xb].Winner
			} else if graph[yb][xb].IsSeed {
				graph[i][j].Winner = graph[ya][xa].Winner
			} else {
				graph[i][j].TeamA.Name = graph[ya][xa].Winner
				graph[i][j].TeamB.Name = graph[yb][xb].Winner
			}
		}
	}

	// トーナメント表の第2層について更新する
	for i := 0; i < len(graph[LayerCount-3]); i++ {
		xa := graph[LayerCount-3][i].TeamA.ChildX
		ya := graph[LayerCount-3][i].TeamA.ChildY
		xb := graph[LayerCount-3][i].TeamB.ChildX
		yb := graph[LayerCount-3][i].TeamB.ChildY
		if graph[ya][xa].Winner != "" && graph[yb][xb].Winner != "" {
			graph[LayerCount-2][i*2].TeamA.Name = graph[ya][xa].Winner
			graph[LayerCount-2][i*2].TeamB.Name = graph[yb][xb].Winner
			graph[LayerCount-2][i*2+1].TeamA.Name = graph[ya][xa].Winner
			graph[LayerCount-2][i*2+1].TeamB.Name = graph[yb][xb].Winner
		}

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
			if graph[ya][xa].MatchNum == 0 {
				xa1 := graph[ya][xa].TeamA.ChildX
				ya1 := graph[ya][xa].TeamA.ChildY
				xa, ya = xa1, ya1
			}
			if graph[yb][xb].MatchNum == 0 {
				xb1 := graph[yb][xb].TeamB.ChildX
				yb1 := graph[yb][xb].TeamB.ChildY
				xb, yb = xb1, yb1
			}
			if graph[ya][xa].IsSeed {
				graph[i][j].Winner = graph[yb][xb].Winner
			} else if graph[yb][xb].IsSeed {
				graph[i][j].Winner = graph[ya][xa].Winner
			} else {
				graph[i][j].TeamA.Name = graph[ya][xa].Winner
				graph[i][j].TeamB.Name = graph[yb][xb].Winner
			}
		}
	}

	return graph
}

// RegisterMatchTime 競技時間を登録する
func RegisterMatchTime(graph [][]Match, times []Time) [][]Match {
	for idx, time := range times {
		for i := 0; i < len(graph)-1; i++ {
			for j := 0; j < len(graph[i]); j++ {
				if graph[i][j].MatchNum == idx+1 {
					graph[i][j].Time = time
				}
			}
		}
	}
	return graph
}
