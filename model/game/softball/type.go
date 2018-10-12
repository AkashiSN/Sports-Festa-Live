package softball

import (
	"github.com/AkashiSN/Sports-Festa-Live/model/matche"
)

type Softball struct {
	Matches []matche.Matche
	Queue   struct {
		courtA string
		courtB string
	}
}
