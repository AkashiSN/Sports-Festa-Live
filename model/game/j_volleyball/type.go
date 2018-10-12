package j_volleyball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type J_volleyball struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
