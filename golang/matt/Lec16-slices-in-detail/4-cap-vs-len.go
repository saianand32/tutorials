package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[0:2]
	c := b[0:4]

	fmt.Println(b) // [1 2]
	fmt.Println(c) // [1 2 3 4]

	//why the above behaviour ? coz the capacity of the sliced slices b and c are same as a that is underlying slice

	//now if u want to cur this behaviour u need to use [0:1:1] kind of slicing
	// where [i:j:k] len = j-i and cap = k-i

	b = a[0:2:2]
	c = b[0:1]
	// d := b[0:4] //this will give error as capacity of b < 4 and is equal to 2 as in line 18

	fmt.Println(b) // [1 2]
	fmt.Println(c) // [1]

	//WTF moment
	a1 := []int{1, 2, 3}
	b1 := a1[:2]
	b1[1] = 99
	fmt.Println(a1)        // [1 99 3]
	fmt.Println(b1)        // [1 99]
	fmt.Printf("%p\n", a1) // 0x140000aa018 [NOTE: for slices u can use %p with a1 else if a1 was array u would use %p with &a1]
	fmt.Printf("%p\n", b1) // 0x140000aa018

	//so you realize that if u modify a slice its parent also gets modified as u can see both addresses are same
	//therefore if you need to create a true copy of a slice in go u need to use [::] and append

	c1 := a1[0:2]
	c1 = append(c1, 33)
	fmt.Println(a1)        // [1 99 33]
	fmt.Println(c1)        // [1 99 33]
	fmt.Printf("%p\n", c1) //0x140000aa018 still same address

	c1 = a1[0:2:2]
	c1 = append(c1, 44)
	fmt.Println(a1)        // [1 99 33]
	fmt.Println(c1)        // [1 99 44]
	fmt.Printf("%p\n", c1) //0x140000ba000 this creates a new address location and now u have a copy of slice and not reference

	// So, ALways use [::] double size slice by specifying cap of the sliced array

}
