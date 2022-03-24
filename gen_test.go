package gen_test

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func addInt(a, b int) int {
	return a + b
}
func addFloat64(a, b float64) float64 {
	return a + b
}
func addString(a, b string) string {
	return a + b
}

func addTypeSwitch(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case float64:
		return a.(float64) + b.(float64)
	case string:
		return a.(string) + b.(string)
	default:
		panic("unreachable")
	}
}

func addReflection(a, b interface{}) interface{} {
	if reflect.TypeOf(a).Kind() == reflect.Int {
		return a.(int) + b.(int)
	}
	if reflect.TypeOf(a).Kind() == reflect.Float64 {
		return a.(float64) + b.(float64)
	}
	if reflect.TypeOf(a).Kind() == reflect.String {
		return a.(string) + b.(string)
	}
	panic("unreachable")
}

func addGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

type typeSet interface {
	int | float64 | string
}

func addGenericsWithTypeSet[T typeSet](a, b T) T {
	return a + b
}

func BenchmarkNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addInt(rand.Int(), rand.Int())
		addFloat64(rand.Float64(), rand.Float64())
		addString(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkTypeSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTypeSwitch(rand.Int(), rand.Int())
		addTypeSwitch(rand.Float64(), rand.Float64())
		addTypeSwitch(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addReflection(rand.Int(), rand.Int())
		addReflection(rand.Float64(), rand.Float64())
		addReflection(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addGenerics(rand.Int(), rand.Int())
		addGenerics(rand.Float64(), rand.Float64())
		addGenerics(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkGenericsWithTypeSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addGenericsWithTypeSet(rand.Int(), rand.Int())
		addGenericsWithTypeSet(rand.Float64(), rand.Float64())
		addGenericsWithTypeSet(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}
