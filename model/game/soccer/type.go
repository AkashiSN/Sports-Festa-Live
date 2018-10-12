package soccer

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Soccer struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
