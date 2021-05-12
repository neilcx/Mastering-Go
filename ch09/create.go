package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}

//➜  ch09 git:(master) ✗ go run create.go -n 9
//Going to create 9 goroutines.
//4 1 2 0 7 5 8 6 3
//Exiting...
//➜  ch09 git:(master) ✗ go run create.go -n 8
//Going to create 8 goroutines.
//0 3 2 5 4 6 7 1
//Exiting...
//➜  ch09 git:(master) ✗ go run create.go -n 100
//Going to create 100 goroutines.
//5 1 2 3 4 25 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 6 42 26 27 28 29 30 8 7 36 32 33 34 35 58 43 44 45 37 47 46 31 48 49 50 51 52 53 54 55 56 57 40 39 41 66 59 60 61 62 63 64 65 70 67 68 69 72 71 73 74 38 87 0 88 80 93 89 90 91 82 81 92 77 78 86 84 85 96 94 95 98 97 99 75 83 79 76
//Exiting...