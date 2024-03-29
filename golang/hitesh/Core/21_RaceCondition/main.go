package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	// mut := &sync.RWMutex{}
	//---- Can also use RWMutex
	// then u can use mut.RLock() mut.RUnlock()
	// then u can use mut.WLock() mut.WUnlock()

	score := []int{0}

	//can add wg.Add(1) 3 times else u can also do this as only 3 go routines

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()

		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()

		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()

		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

//To ckeck if prog has race run $ go run --race
