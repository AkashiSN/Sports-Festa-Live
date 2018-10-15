package match

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// LoadTime 競技時間を読み込む
func LoadTime() Time {
	bytes, err := ioutil.ReadFile("conf/match.json")
	if err != nil {
		log.Fatal(err)
	}
	var time Time
	if err := json.Unmarshal(bytes, &time); err != nil {
		log.Fatal(err)
	}
	return time
}
