package badminton

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Badminton struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
