package team

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// LoadTeam チーム情報を読み込む
func LoadTeam() Team {
	bytes, err := ioutil.ReadFile("conf/team.json")
	if err != nil {
		log.Fatal(err)
	}
	var team Team
	if err := json.Unmarshal(bytes, &team); err != nil {
		log.Fatal(err)
	}
	return team
}
