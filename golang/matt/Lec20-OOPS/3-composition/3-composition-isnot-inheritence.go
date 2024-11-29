package main

import "fmt"

type Pair struct {
	len     int
	breadth int
}

type PairWithHeight struct {
	Pair   // Embedded type (composition)
	height int
}

func f1(x Pair) int {
	return x.len
}

func main() {
	a := Pair{5, 10}
	b := PairWithHeight{Pair{5, 10}, 20}

	fmt.Println(f1(a)) // 5
	// fmt.Println(f1(b)) // error

	// f1(b) gives an error because `PairWithHeight` and `Pair` are distinct concrete types,
	// even though `PairWithHeight` contains a `Pair` field.
	// In Go, embedding one type into another does not create inheritance (like in some other languages).
	// It creates a "has-a" relationship, not an "is-a" relationship. Therefore, even though
	// `PairWithHeight` embeds `Pair`, the function `f1` expects a `Pair`, not a `PairWithHeight`.
	//
	// To fix this, we could either:
	// 1. Pass only the `Pair` part of `b` (b.Pair) to the function, or
	// 2. Modify `f1` to accept a `PairWithHeight` (if applicable).
	// fmt.Println(f1(b)) // This would give an error, as `PairWithHeight` is not assignable to `Pair`.

	// to get around this u must smartly use an interface (u know how)
}
