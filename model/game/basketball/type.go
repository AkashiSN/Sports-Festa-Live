package basketball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Basketball struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}