package match

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// LoadTime 競技時間を読み込む
func LoadTime() Times {
	bytes, err := ioutil.ReadFile("conf/match.json")
	if err != nil {
		log.Fatal(err)
	}
	var time Times
	if err := json.Unmarshal(bytes, &time); err != nil {
		log.Fatal(err)
	}
	return time
}

// DebugDump デバックように出力する
func DebugDump(graph [][]Match) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			fmt.Printf("\nmatchNum:%d\n", graph[i][j].MatchNum)
			fmt.Printf("parent(x,y):(%d,%d)\n", j, i)
			fmt.Printf("childA(x,y):(%d,%d)\n", graph[i][j].TeamA.ChildX, graph[i][j].TeamA.ChildY)
			fmt.Printf("childB(x,y):(%d,%d)\n", graph[i][j].TeamB.ChildX, graph[i][j].TeamB.ChildY)
			fmt.Printf("%s VS %s\n", graph[i][j].TeamA.Name, graph[i][j].TeamB.Name)
			fmt.Printf("Winner:%s\n", graph[i][j].Winner)
		}
	}
}
