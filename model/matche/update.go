package matche

// UpdateFirstLayer トーナメント表の第1層の計算を行う
func UpdateFirstLayer(teams []string) []Matche {
	matchNum := len(teams) - 1
	matches := make([]Matche, matchNum)
	matchCount := 0

	for i := 0; i < matchNum; i += 4 {
		for j := i; j < i+3; j += 1 {
			if teams[j] == "" && teams[j+1] == "" {
				matches[matchCount].TeamA = teams[j-1]
				matches[matchCount].TeamB = teams[j+2]
				matches[matchCount].MatcheNum = matchCount
				matchCount += 1
			}
			if j != i+1 && teams[j] != "" && teams[j+1] != "" {
				matches[matchCount].TeamA = teams[j]
				matches[matchCount].TeamB = teams[j+1]
				matches[matchCount].MatcheNum = matchCount
				matchCount += 1
			}

		}
	}
	return matches
}

// UpdateSecondLayer トーナメント表の第2層の計算を行う
func UpdateSecondLayer(teams []string, matches []Matche) []Matche {
	matchNum := len(teams) - 1
	matchCount := 0

	for i := 0; i < matchNum; i += 1 {
		if matches[i].TeamA != "" {
			matchCount += 1
		}
	}

	for i := 0; i < matchNum; i += 4 {
		if (teams[i+1] == "") != (teams[i+2] == "") {
			if teams[i+1] == "" {
				matches[matchCount].MatcheNum = matchCount
				matches[matchCount].TeamA = teams[i]
				matches[matchCount].TeamB = matches[i/4].Winner
				matchCount += 1
			}
			if teams[i+2] == "" {
				matches[matchCount].MatcheNum = matchCount
				matches[matchCount].TeamA = matches[i/4].Winner
				matches[matchCount].TeamB = teams[i+3]
				matchCount += 1
			}
		} else {
			matches[matchCount].MatcheNum = matchCount
			matches[matchCount].TeamA = matches[i/4].Winner
			matches[matchCount].TeamB = matches[i/4+2].Winner
			matchCount += 1
		}
	}
	return matches
}
