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
