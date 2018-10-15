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
		graph[i] = make([]Match, int(math.Pow(2, float64(i+1))))
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

	// 試合番号を振る
	matchCount := 0
	// トーナメント表の第1層について
	for i := 0; i < int(math.Pow(2, float64(LayerCount))); i += 4 {
		for j := i; j < i+3; j++ {
			if teams[j] == "" && teams[j+1] == "" {
				graph[LayerCount-1][i].MatchNum = matchCount
				matchCount++
			}
			if j != i+1 && teams[j] != "" && teams[j+1] != "" {
				graph[LayerCount-1][i].MatchNum = matchCount
				matchCount++
			}
		}
	}
	// トーナメント表の第2層について
	count := 0
	for i := 0; i < int(math.Pow(2, float64(LayerCount))); i += 4 {
		if (teams[i+1] == "") != (teams[i+2] == "") {
			if teams[i+1] == "" {
				graph[LayerCount-2][count].MatchNum = matchCount
				matchCount++
			}
			if teams[i+2] == "" {
				graph[LayerCount-2][count].MatchNum = matchCount
				matchCount++
			}
		} else if !(teams[i+1] == "" && teams[i+2] == "") {
			graph[LayerCount-2][count].MatchNum = matchCount
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
