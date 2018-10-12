package team

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadTeam() Team {
	bytes, err := ioutil.ReadFile("team.json")
	if err != nil {
		log.Fatal(err)
	}
	var team Team
	if err := json.Unmarshal(bytes, &team); err != nil {
		log.Fatal(err)
	}
	return team
}
