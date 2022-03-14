package main

import (
	"fmt"
	"github.com/FernandoDucha/sawwebservice"
	"reflect"
)

func main() {
	fooType := reflect.TypeOf(sawwebservice.PointsVector{})
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
}
