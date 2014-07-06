package reflectutils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func _to_json(x interface{}) string {
	p, _ := json.MarshalIndent(x, "", "  ")
	return string(p)
}

type CTest struct {
	C_int64 int64
}

type DTest struct {
	D_int64 int64
}

type ATest struct {
	CTest
	*DTest
	A       string
	B       int
	C       *string
	D       *int
	A_int64 int64
}

type BTest struct {
	*ATest
	BA   string
	BB   int
	BC   *string
	BD   *int
	Time time.Time
}

func Test_ToMap(t *testing.T) {
	a, b, c, d := "A", 1, "C", 4

	obj := &BTest{ATest: &ATest{A: a, B: b, C: &c, D: &d}}
	obj.C_int64 = 7
	obj.Time = time.Now()

	m := ToMap(obj)

	for _, key := range []string{"BA", "BB", "BC", "BD", "Time", "C_int64", "A", "B", "C", "D", "A_int64"} {
		if _, ok := m[key]; !ok {
			t.Fatal("not found", key)
		}
	}

	if _, ok := m["D_int64"]; ok {
		t.Fatal("D_int64 must not be include in map, cause structure containing it is nil")
	}

	test_map_val(m, "Time", fmt.Sprint(obj.Time), t)
	test_map_val(m, "BA", fmt.Sprint(obj.BA), t)
	test_map_val(m, "BB", fmt.Sprint(obj.BB), t)
	test_map_val(m, "BC", fmt.Sprint(obj.BC), t)
	test_map_val(m, "BD", fmt.Sprint(obj.BD), t)

	test_map_val(m, "A", fmt.Sprint(obj.A), t)
	test_map_val(m, "B", fmt.Sprint(obj.B), t)
	test_map_val(m, "C", fmt.Sprint(c), t)
	test_map_val(m, "D", fmt.Sprint(d), t)

	test_map_val(m, "C_int64", fmt.Sprint(obj.C_int64), t)
}

func test_map_val(m map[string]interface{}, key, expect string, t *testing.T) {
	val, ok := m[key]
	if !ok {
		t.Fatal(key)
		return
	}

	if fmt.Sprint(val) != expect {
		t.Fatal("values not match", key, expect, val)
	}
}

func Test_ToMap_Alias(t *testing.T) {
	type x map[string]interface{}

	test := make(x)
	test["x"] = 1

	m := ToMap(test)

	if m["x"] != 1 {
		t.Fail()
	}
}

func TestForeach(t *testing.T) {
	slice := []int{1, 2, 3, 4}

	Foreach(slice, func(i int, val interface{}) bool {
		t.Log(i, val)
		return true
	})
}

func TestIsSlice(t *testing.T) {
	if !IsSlice([]int{0}) {
		t.Fail()
	}

	var slice []int

	if !IsSlice(slice) {
		t.Fail()
	}

	if IsSlice(1) {
		t.Fail()
	}
}

func _foreach(arr []bool, fn func(int, bool) bool) {
	for i, v := range arr {
		if !fn(i, v) {
			break
		}
	}
}

func BenchmarkRange(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]bool, 16)
		b.StartTimer()
		_foreach(slice, func(i int, val bool) bool {
			if i == 0 && val {

			}
			return true
		})
		b.StopTimer()
	}
}

func BenchmarkRangeTillFirst(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]bool, 16)
		b.StartTimer()
		_foreach(slice, func(i int, val bool) bool {
			return false
		})
		b.StopTimer()
	}
}

func BenchmarkForeach(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]bool, 16)
		b.StartTimer()
		Foreach(slice, func(i int, val interface{}) bool {
			if i == 0 {

			}
			return true
		})
		b.StopTimer()
	}
}

func BenchmarkForeachTillFirst(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]bool, 16)
		b.StartTimer()
		Foreach(slice, func(i int, val interface{}) bool {
			return false
		})
		b.StopTimer()
	}
}

func BenchmarkToMap(b *testing.B) {
	b.StopTimer()
	obj := BTest{}
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		ToMap(obj)
		b.StopTimer()
	}
}

func BenchmarkToMapNocahe(b *testing.B) {
	b.StopTimer()
	cache_layouts = false
	get_layout_cache = make(map[reflect.Type]type_layout)
	obj := BTest{}
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		ToMap(obj)
		b.StopTimer()
	}
	cache_layouts = true
}
