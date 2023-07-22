package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				meter.add()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.value())

	// go run basic/sync_race_condition.go
	// 713393
	// go run basic/sync_race_condition.go
	// 710762
	// go run basic/sync_race_condition.go
	// 724770

	/* BEFORE
	go run --race basic/sync_race_condition.go
	==================
	WARNING: DATA RACE
	Read at 0x00c0000180b8 by goroutine 16:
	  main.(*counter).add()
	      /Users/doddyharyadi/melabs/go/bolang/basic/sync_race_condition.go:43 +0x48
	  main.main.func1()
	      /Users/doddyharyadi/melabs/go/bolang/basic/sync_race_condition.go:20 +0x44

	Previous write at 0x00c0000180b8 by goroutine 6:
	  main.(*counter).add()
	      /Users/doddyharyadi/melabs/go/bolang/basic/sync_race_condition.go:43 +0x58
	  main.main.func1()
	      /Users/doddyharyadi/melabs/go/bolang/basic/sync_race_condition.go:20 +0x44

	Goroutine 16 (running) created at:
	  main.main()
	      /Users/doddyharyadi/melabs/go/bolang/basic/sync_race_condition.go:18 +0x70

	Goroutine 6 (finished) created at:
	  main.main()
	      /Users/doddyharyadi/melabs/go/bolang/basic/sync_race_condition.go:18 +0x70
	==================
	708788
	Found 1 data race(s)
	exit status 66
	*/

	/* AFTER add sync.Mutex
	go run --race basic/sync_race_condition.go
	1000000
	*/
}

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) add() {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) value() int {
	return c.val
}
