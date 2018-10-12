package tennis

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Tennis struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
