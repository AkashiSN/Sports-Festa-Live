package tabletennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Tabletennis struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
