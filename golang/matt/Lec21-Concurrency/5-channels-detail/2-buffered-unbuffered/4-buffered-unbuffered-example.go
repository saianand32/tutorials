package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	t bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t
	t.t = true // UNSAFE AT ANY SPEED
}

func main() {
	vs, ch := make([]T, 5), make(chan *T)
	for i := range vs {
		go send(i, ch)
	}
	time.Sleep(1 * time.Second) // all goroutines started
	for i := range vs {         // copy quickly!
		vs[i] = *<-ch
	}
	for _, v := range vs { // print later
		fmt.Println(v)
	}
}

// output: all false as recieve finished before send finishes ****
// // {1 false}
// // {2 false}
// // {4 false}
// // {3 false}
// // {5 false}

// in abive in line 20 if u change it to buffered make(chan *T, 5)
// output will be all true as send finishes off before itself as it stores in buffer **
// // {1 true}
// // {4 true}
// // {2 true}
// // {3 true}
// // {5 true}
