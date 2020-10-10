package main

import (
	"fmt"
	"testing"
)

type intersection interface {
	Intersection(int, int, []int, []int) float32
}

// TestItem type
type TestItem struct {
	inputs   Args
	result   float32
	hasError bool
}

// Args ...
type Args struct {
	m  int
	n  int
	v1 []int
	v2 []int
}

func TestResultInt(t *testing.T) {
	var m = []TestItem{
		{
			Args{
				4, 6, []int{1, 3, 2, 6}, []int{7, 0, 9, 4, 3, 1},
			},
			2, false,
		}, {
			Args{
				4, 6, []int{19, 5, 2, 6}, []int{5, 0, 9, 4, 18, 56},
			},
			5, false,
		}, {
			Args{
				2, 4, []int{2, 4}, []int{2, 3, 4, 5},
			},
			3, false,
		}, {
			Args{
				0, 4, []int{}, []int{2, 3, 4, 5},
			},
			0, false,
		}, {
			Args{
				2, 0, []int{2, 4}, []int{},
			},
			0, false,
		}, {
			Args{
				4, 4, []int{1, 7, 8, 9}, []int{2, 3, 4, 5},
			},
			0, false,
		},
	}
	for _, i := range m {
		testExecute(i, t)
	}
	fmt.Println("\n-> Int results: TestResultInt()")
	fmt.Printf("Tests done: %v \n", len(m))
}

func TestNegativeErrorNM(t *testing.T) {
	var m = []TestItem{
		{
			Args{
				-1, 4, []int{}, []int{2, 3, 4, 5},
			},
			-1, true,
		}, {
			Args{
				-2, 4, []int{3, 4}, []int{2, 3, 4, 5},
			},
			-1, true,
		}, {
			Args{
				-1, -2, []int{}, []int{},
			},
			-1, true,
		}, {
			Args{
				-3, -2, []int{3, 4, 8}, []int{2, 3},
			},
			-1, true,
		},
	}
	for _, i := range m {
		testExecute(i, t)
	}
	fmt.Println("\n-> Negative 'n' or 'm': TestNegativeErrorNM()")
	fmt.Printf("Tests done: %v \n", len(m))
}

func TestNegativeErrorVx(t *testing.T) {
	var m = []TestItem{
		{
			Args{
				1, 4, []int{2}, []int{2, 3, -4, 5},
			},
			-1, true,
		}, {
			Args{
				2, 4, []int{-3, 4}, []int{2, -3, 4, 5},
			},
			-1, true,
		}, {
			Args{
				4, 2, []int{4, 2, 0, -1}, []int{3, -1},
			},
			-1, true,
		}, {
			Args{
				8, 9, []int{3, 4, 8, 9, 7, 2, 6, -1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			-1, true,
		},
	}
	for _, i := range m {
		testExecute(i, t)
	}
	fmt.Println("\n-> Negative int inside vector: TestNegativeErrorVx()")
	fmt.Printf("Tests done: %v \n", len(m))
}

func TestZeroDivision(t *testing.T) {
	var m = []TestItem{
		{
			Args{
				2, 4, []int{2, 0}, []int{4, 3, 0, 5},
			},
			0, false,
		}, {
			Args{
				5, 4, []int{3, 4, 8, 1, 0}, []int{6, 2, 7, 0},
			},
			0, false,
		},
	}
	for _, i := range m {
		testExecute(i, t)
	}
	fmt.Println("\n-> Results made with '0/int' operation: TestZeroDivision()", len(m))
	fmt.Printf("Tests done: %v \n", len(m))
	fmt.Println()
}

func TestDifferentLengthError(t *testing.T) {
	var m = []TestItem{
		{
			Args{
				5, 4, []int{2, 0}, []int{4, 3, 0, 5},
			},
			-1, true,
		}, {
			Args{
				2, 4, []int{3, 4, 8, 1, 0}, []int{6, 2, 7, 0},
			},
			-1, true,
		}, {
			Args{
				2, 7, []int{2, 0}, []int{4, 3, 0, 5},
			},
			-1, true,
		}, {
			Args{
				2, 2, []int{2, 0}, []int{4, 3, 0, 5},
			},
			-1, true,
		}, {
			Args{
				5, 8, []int{2, 0}, []int{4, 3, 0, 5},
			},
			-1, true,
		}, {
			Args{
				2, 3, []int{3, 4, 8, 1, 0}, []int{6, 2, 7, 0},
			},
			-1, true,
		},
  }
  
	for _, i := range m {
		testExecute(i, t)
	}
	fmt.Println("\n-> Results made with '0/int' operation: TestZeroDivision()", len(m))
	fmt.Printf("Tests done: %v \n", len(m))
	fmt.Println()
}

func BenchmarkIntersection(b *testing.B) {
	m := 4
	n := 6
	v1 := []int{3, 2, 1}
	v2 := []int{9, 8, 0, 2, 1}

	for i := 0; i < b.N; i++ {
		Intersection(m, n, v1, v2)
	}
}

func testExecute(item TestItem, t *testing.T) {
	// get result of func()
	result, err := Intersection(item.inputs.m, item.inputs.n, item.inputs.v1, item.inputs.v2)

	if item.hasError {
		// expected an error
		if err == nil {
			t.Errorf("\nFAILED!\n☐ Intersection()\n  %v \n- Expected: an error,\n✘ Got: value '%v'", item.inputs, result)
		} else {
			t.Logf("\nPASSED!\n☐ Intersection()\n  %v \n- Expected: an error,\n✔ Got: an error '%v'\n", item.inputs, err)
		}

	} else {
		// expected a value
		if result != item.result {
			t.Errorf("\nFAILED!\n☐ Intersection()\n  %v \n- Expected: %v,\n✘ Got: %v \n", item.inputs, item.result, result)
		} else {
			t.Logf("\nPASSED!\n☐ Intersection() \n  %v \n- Expected: %v \n✔ Got: %v \n", item.inputs, item.result, result)
		}
	}
}
