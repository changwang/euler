package challenge29

import (
	"math/big"
)

func DistinctPowers(coeff int64) int64 {
	var i, j int64
	m := big.NewInt(0)

	distincPows := make(map[string]bool, 0)
	for i = 2; i <= coeff; i++ {
		for j = 2; j <= coeff; j++ {
			pow := big.NewInt(1)
			pow = pow.Exp(big.NewInt(i), big.NewInt(j), m)
			if !distincPows[pow.String()] {
				distincPows[pow.String()] = true
			}
		}
	}
	return int64(len(distincPows))
}
