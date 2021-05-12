package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	//waitGroup.Add(1)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		fmt.Printf("%#v\n", waitGroup)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	//waitGroup.Done()
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}

//➜  ch09 git:(master) ✗ go run syncGo.go -n 10
//Going to create 10 goroutines.
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x0, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x1, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x2, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x3, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x4, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x5, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x6, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x7, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x8, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x9, 0x0}}
//0 4 2 3 5 6 sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x5, 0x0}}
//sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x3, 0x0}}
//9 8 7 1
//Exiting...
//panic: sync: negative WaitGroup counter
//
//goroutine 19 [running]:
//sync.(*WaitGroup).Add(0xc00008c030, 0xffffffffffffffff)
///usr/local/Cellar/go/1.13.5/libexec/src/sync/waitgroup.go:74 +0x139
//sync.(*WaitGroup).Done(0xc00008c030)
///usr/local/Cellar/go/1.13.5/libexec/src/sync/waitgroup.go:99 +0x34
//main.main.func1(0xc00008c030, 0x1)
///Users/chenxi1/workspace/projects/Mastering-Go/ch09/syncGo.go:24 +0xf2
//created by main.main
///Users/chenxi1/workspace/projects/Mastering-Go/ch09/syncGo.go:21 +0x30b
//exit status 2
//➜  ch09 git:(master) ✗