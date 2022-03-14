package pointgenerator

import (
	"math/rand"
	"sawwebservice/sawwebservice"
)

func pointgen() sawwebservice.PointsVector {
	pv, _ := sawwebservice.NewPointsVector(2)
	rand.Seed(23498767)
	var ps []sawwebservice.Point
	for i := 0; i < 100; i++ {
		p := make(sawwebservice.Point, 2)
		for j := 0; j < 2; j++ {
			p[j] = rand.Float64() * 100
		}
		ps = append(ps, p)
		// if inserted{

		// }
	}
	pv.PlacePoints(ps)
	return pv
}
