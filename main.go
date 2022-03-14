package pointgenerator

import (
	"math/rand"

	"github.com/FernandoDucha/sawwebservice"
)

func pointgen() {
	pv, _ := sawwebservice.NewPointsVector(2)
	rand.Seed(23498767)
	for i := 0; i < 100; i++ {
		p := make([]float64, 2)
		for j := 0; j < 2; j++ {
			p[i] = rand.Float64() * 100
		}
		inserted, _ := pv.insertPoint(p)
	}

}