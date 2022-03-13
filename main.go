package main

import (
	"math/rand"
	"sawwebservice"
)

func main() {
	pv := sawwebservice.NewPointsVector(2)
	rand.Seed(23498767)
	for i := 0; i < 100; i++ {
		p := make(sawwebservice.Point, 2)
		for j := 0; j < 2; j++ {
			p[i] = rand.Float64() * 100
		}
		pv.insertPoint(p)
	}
}
