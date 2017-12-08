package main

import (
	"reflect"
	"time"
	"regexp"
	"strconv"
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

type Res2 struct {
	Failed int64
	Passed int64
	Ignored int64
}


//Write a function that takes a string and parses it to struct
//type tests struct {
//	failed  int64
//	passed  int64
//	ignored int64
//}
//The function should return the tests struct.
func func2(str string) Res2 {
	result := Res2{}
	re0 := regexp.MustCompile(`^Tests`)
	str0 := re0.FindAllString(str, -1)
	if len(str0) == 0 {
		return result
	}
	re := regexp.MustCompile(`(failed|passed|ignored|muted):\s(\d+)`)
	submatch := re.FindAllStringSubmatch(str, -1)
	for _, row := range submatch {
		if len(row) < 3 {
			continue
		}
		switch row[1] {
		case "failed":
			val, _ := strconv.Atoi(row[2])
			result.Failed = int64(val)
			break
		case "passed":
			val, _ := strconv.Atoi(row[2])
			result.Passed = int64(val)
			break
		case "ignored":
			val, _ := strconv.Atoi(row[2])
			result.Ignored += int64(val)
			break
		case "muted":
			val, _ := strconv.Atoi(row[2])
			result.Ignored += int64(val)
			break
		}
	}
	return result
}
