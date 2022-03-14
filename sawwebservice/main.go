package main

import (
	"fmt"
	"reflect"
	"sawwebservice/sawwebservice"
)

type pointvec sawwebservice.PointsVector

func main() {
	// pv, _ := sawwebservice.NewPointsVector(2)
	fooType := reflect.TypeOf(pointvec{})
	fmt.Printf("%d\n", fooType.NumMethod())
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Printf("%d %s \n", i, method.Name)
	}
}
