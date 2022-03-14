package main

import (
	"fmt"
	"reflect"

	"github.com/FernandoDucha/sawwebservice"
)

func main() {
	pv, _ := sawwebservice.NewPointsVector(2)
	fooType := reflect.TypeOf(pv)
	fmt.Printf("%d\n", fooType.NumMethod())
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Printf("%d %s \n", i, method.Name)
	}
}
