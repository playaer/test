package main

import (
	"reflect"
	"time"
)

func main() {

}

// Write a function that takes a one required parameter int and three optional:
// first - int, default value 10
// second - time.Time, default time.Now()
// third - pointer to any struct, default - nil
// The function should return all passed parameters.

// !!! нет нормальной проверки на тип переменных
func func1(param1 int, params ...interface{}) (int, int, time.Time, *struct{}) {
	result := []interface{}{}
	result = append(result, param1)

	var p1 int = 10
	var p2 time.Time = time.Now()
	var p3 *struct{} = nil

	if len(params) > 0 && reflect.TypeOf(params[0]) == reflect.TypeOf(p1) {
		p1 = params[0].(int)
	}

	if len(params) > 1 && reflect.TypeOf(params[1]) == reflect.TypeOf(p2) {
		p2 = params[1].(time.Time)
	}

	if len(params) > 2 && reflect.TypeOf(params[2]) == reflect.TypeOf(p3) {
		p3 = params[2].(*struct{})
	}

	return param1, p1, p2, p3
}
