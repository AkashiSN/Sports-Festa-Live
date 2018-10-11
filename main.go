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

type Matches struct {
	TeamA  string
	TeamB  string
	Winner string
}

func main() {
	bytes, err := ioutil.ReadFile("team.json")
	if err != nil {
		log.Fatal(err)
	}
	var team Team
	if err := json.Unmarshal(bytes, &team); err != nil {
		log.Fatal(err)
	}

	matchNum := len(team.Dodge) - 1
	matches := make([]Matches, matchNum)
	matchCount := 0

	for i := 0; i < len(team.Dodge); i += 4 {
		for j := i; j < i+3; j += 1 {
			if team.Dodge[j] == "" && team.Dodge[j+1] == "" {
				matches[matchCount].TeamA = team.Dodge[j-1]
				matches[matchCount].TeamB = team.Dodge[j+2]
				matchCount += 1
			}
			if team.Dodge[j] != "" && team.Dodge[j+1] != "" {
				matches[matchCount].TeamA = team.Dodge[j]
				matches[matchCount].TeamB = team.Dodge[j+1]
				matchCount += 1
			}
		}
	}

	for i := 0; i < len(team.Dodge); i += 4 {
		if (team.Dodge[i+1] == "") != (team.Dodge[i+2] == "") {
			if team.Dodge[i+1] == "" {
				matches[matchCount].TeamA = team.Dodge[i]
				matches[matchCount].TeamB = matches[i/4].Winner
				matchCount += 1
			}
			if team.Dodge[i+2] == "" {
				matches[matchCount].TeamA = matches[i/4].Winner
				matches[matchCount].TeamB = team.Dodge[i+3]
				matchCount += 1
			}
		}
	}

	fmt.Print(matches)
}
