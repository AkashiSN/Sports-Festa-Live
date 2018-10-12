package j_dodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type J_dodge struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
