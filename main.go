package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Team struct {
	Dodge []string `json:"dodge"`
}

type Matche struct {
	TeamA  string
	TeamB  string
	Winner string
}

type Dodge struct {
	matche     []Matche
	matchCount int
}

func initMatch(teams []string) (int, []Matche) {
	matchNum := len(teams) - 1
	matches := make([]Matche, matchNum)
	matchCount := 0

	for i := 0; i < len(teams); i += 4 {
		for j := i; j < i+3; j += 1 {
			if teams[j] == "" && teams[j+1] == "" {
				matches[matchCount].TeamA = teams[j-1]
				matches[matchCount].TeamB = teams[j+2]
				matchCount += 1
			}
			if teams[j] != "" && teams[j+1] != "" {
				matches[matchCount].TeamA = teams[j]
				matches[matchCount].TeamB = teams[j+1]
				matchCount += 1
			}
		}
	}
	return matchCount, matches
}

func loadJson() Team {
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

func updateSecondLayer(teams []string, matchCount int, matches []Matche) []Matche {
	for i := 0; i < len(teams); i += 4 {
		if (teams[i+1] == "") != (teams[i+2] == "") {
			if teams[i+1] == "" {
				matches[matchCount].TeamA = teams[i]
				matches[matchCount].TeamB = matches[i/4].Winner
				matchCount += 1
			}
			if teams[i+2] == "" {
				matches[matchCount].TeamA = matches[i/4].Winner
				matches[matchCount].TeamB = teams[i+3]
				matchCount += 1
			}
		}
	}
	return matches
}

func main() {
	team := loadJson()
	var dodge Dodge
	dodge.matchCount, dodge.matche = initMatch(team.Dodge)
	dodge.matche = updateSecondLayer(team.Dodge, dodge.matchCount, dodge.matche)
	fmt.Print(dodge.matche)
}
