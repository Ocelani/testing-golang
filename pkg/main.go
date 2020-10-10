package main

import (
	"fmt"
	_ "net/http/pprof"
)

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
// var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

type negativeIntError struct {
	e bool
}

func (e *negativeIntError) Error() string {
	return fmt.Sprintf("%v", e.e)
}

// Intersection compares the values between v1[] and v2[],
// then, calculates the mean of the intersection numbers.
func Intersection(m int, n int, v1 []int, v2 []int) (r float32, err error) {
	if err := (m != len(v1) || n != len(v2)); err == true {
		err := &negativeIntError{err}
		return -1, err
	}

	sum := 0
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if err := (v1[i] < 0 || v2[j] < 0); err == true {
				err := &negativeIntError{err}
				return -1, err
			}
			if v1[i] == v2[j] {
				sum += v1[i]
				count++
			}
		}
	}
	if count < 1 {
		return 0, nil
	}
	return float32(sum / count), nil
}

// // CreateVector creates an array (vector)
// // with provided length value.
// func CreateVector(mn int) []int {
// 	r := rand.Perm(20)
// 	v := make([]int, mn)

// 	for i := range v {
// 		v[i] = r[i]
// 	}
// 	return v
// }

// func run() (err error) {
// 	var m, n int
// 	var v1, v2 []int

// 	rand.Seed(time.Now().UnixNano())

// 	m = rand.Intn(10)
// 	n = rand.Intn(10)

// 	cv := make(chan []int, 2)
// 	cv <- CreateVector(m)
// 	cv <- CreateVector(n)
// 	v1, v2 = <-cv, <-cv
// 	close(cv)

// 	printVars(m, n, v1, v2)

// 	r, _ := Intersection(m, n, v1, v2)
// 	printResult(r)

// 	return
// }

// func printVars(m int, n int, v1 []int, v2 []int) {
// 	fmt.Println("------------------------------")
// 	fmt.Println("m:", m)
// 	fmt.Println("v1:", v1)
// 	fmt.Println("\nn:", n)
// 	fmt.Println("v2:", v2)
// }

// func printResult(r float32) {
// 	fmt.Println("\n------------------------------")
// 	fmt.Println("✨ Result:", r)
// 	fmt.Println("------------------------------")
// }

// func main() {
// 	flag.Parse()

// 	if *cpuprofile != "" {
// 		f, err := os.Create(*cpuprofile)
// 		if err != nil {
// 			log.Fatal("could not create CPU profile: ", err)
// 		}
// 		defer f.Close() // error handling omitted for example
// 		if err := pprof.StartCPUProfile(f); err != nil {
// 			log.Fatal("could not start CPU profile: ", err)
// 		}
// 		defer pprof.StopCPUProfile()
// 	}

// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}

// 	if *memprofile != "" {
// 		f, err := os.Create(*memprofile)
// 		if err != nil {
// 			log.Fatal("could not create memory profile: ", err)
// 		}
// 		defer f.Close() // error handling omitted for example
// 		runtime.GC()    // get up-to-date statistics
// 		if err := pprof.WriteHeapProfile(f); err != nil {
// 			log.Fatal("could not write memory profile: ", err)
// 		}
// 	}

// }
