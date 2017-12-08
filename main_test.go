package main

import (
	"reflect"
	"testing"
	"time"
)

var func1Tests = []struct {
	A0 int
	A1 int
	A2 time.Time
	A3 *struct{}
}{
	{5, 10, time.Now(), nil},
	{5, 5, time.Now(), nil},
	{5, 5, time.Now().Add(time.Minute), nil},
	{5, 5, time.Now().Add(time.Minute), &struct{}{}},
}

func TestFunc1(t *testing.T) {
	for i, row := range func1Tests {
		v := reflect.ValueOf(&row).Elem()

		args := []interface{}{}
		for j := 1; j < i+1; j++ {
			args = append(args, v.Field(j).Interface())
		}

		p0, p1, p2, p3 := func1(row.A0, args...)
		// можно заморочиться с циклом здесь, но тогда придется писать if для time
		if row.A0 != p0 {
			t.Error("Test", i, ": expected:", p0, "actual:", row.A0)
		}
		if row.A1 != p1 {
			t.Error("Test", i, ": expected:", p1, "actual:", row.A1)
		}
		if row.A2.Unix() != p2.Unix() {
			t.Error("Test", i, ": expected:", p2, "actual:", row.A2)
		}
		if row.A3 != p3 {
			t.Error("Test", i, ": expected:", p0, "actual:", row.A3)
		}
	}
}

//type Res2 struct {
//	Failed int64
//	Passed int64
//	Ignored int64
//}

var func2Tests = []struct {
	Str string
	Failed int64
	Passed int64
	Ignored int64
}{
	{"Tests passed: 64", 0, 64, 0},
	{"Tests failed: 2 (1 new), passed: 173, ignored: 6, muted: 3; process exited with code 1", 2, 173, 9},
	{"Tests failed: 4 (3 new), passed: 55; process exited with code 1", 4, 55, 0},
	{"Tests failed: 5 (1 new), passed: 125; process exited with code 1", 5, 125, 0},
	{"Tests failed: 1 (1 new), passed: 310, ignored: 13; process exited with code 1", 1, 310, 13},
	{"Tests passed: 311, ignored: 13", 0, 311, 13},
	{"Tests failed: 1 (1 new), passed: 0; process exited with code 1", 1, 0, 0},
	{"Tests failed: 3 (2 new), passed: 174, ignored: 7, muted: 1; process exited with code 1", 3, 174, 8},
	{"Tests passed: 2, failed: 173, muted: 6, ignored: 3; process exited with code 1", 173, 2, 9},
	{"Tests failed: 15, passed: 151; process exited with code 1", 15, 151, 0},
	{"Success", 0, 0, 0},
	{"Process exited with code 1", 0, 0, 0},
	{"Canceled", 0, 0, 0},
	{"Snapshot dependency builds failed: 1", 0, 0, 0},
	{"Canceled (Process exited with code 137)", 0, 0, 0},
}

func TestFunc2(t *testing.T) {
	for i, row := range func2Tests {
		result := func2(row.Str)
		// можно сделать с циклом здесь, но тогда придется писать if для time
		if row.Failed != result.Failed || row.Passed != result.Passed || row.Ignored != result.Ignored {
			t.Error("Test", i, row.Str, ":\n expected:", row.Failed, row.Passed, row.Ignored, "actual:", result)
		}
	}
}
