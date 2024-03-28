package main

import (
	"fmt"
	"sync"
)

var wb sync.WaitGroup

var mut sync.Mutex

func main() {
	mem := 0

	for i := 0; i < 10000; i++ {
		wb.Add(1)
		go addToMemory(&mem)
	}

	wb.Wait()
	fmt.Println(mem)
}

func addToMemory(mem *int) {
	defer wb.Done()

	//need to locak and unlock critical section else u ill get different answers each time
	mut.Lock()
	*mem += 1
	mut.Unlock()
}
