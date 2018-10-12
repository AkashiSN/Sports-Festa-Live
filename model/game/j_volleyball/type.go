package j_volleyball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Manager struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
