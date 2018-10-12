package dodge

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Dodge struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
