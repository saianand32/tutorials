package main

import (
	"fmt"
	"slices"
)

func main() {
	//1. Insert (inserts at index and pushes others to right)
	arr := []string{"Alice", "Bob", "Vera"}
	arr = slices.Insert(arr, 1, "Bill", "Billie")
	arr = slices.Insert(arr, len(arr), "Zac") //insert at end
	fmt.Println(arr)                          //[Alice Bill Billie Bob Vera Zac]

	// 2. Reverse
	arr = []string{"sai", "anand", "utsav", "ankur"}

	slices.Reverse(arr)
	fmt.Println(arr) //[ankur utsav anand sai]

	// 3. Sort

	slices.Sort(arr)
	fmt.Println(arr) //[anand ankur sai utsav]

	//4. SortFunc
	arr2 := []int{1, 2, 3}

	slices.SortFunc(arr2, func(a, b int) int { //unstable sort
		return b - a //descending order
		// return a-b //ascending order
	})

	slices.SortStableFunc(arr2, func(a, b int) int { //stable sort
		return b - a //descending order
		// return a-b //ascending order
	})

	// 5. IsSorted
	fmt.Println(slices.IsSorted(arr2)) //true

	fmt.Println(arr2) //[3 2 1]

	// 6. Min and Max
	fmt.Println(slices.Min(arr2)) //1
	fmt.Println(slices.Max(arr2)) //3

	// 7. index - use it to search for an element, there is no LastIndex method
	arr2 = []int{0, 42, 8}
	fmt.Println(slices.Index(arr2, 8)) //2
	fmt.Println(slices.Index(arr2, 7)) //-1

	// 8. Contains
	fmt.Println(slices.Contains(arr2, 42)) //true

	// 9. BinarySearch
	arr = []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(arr, "Vera")
	fmt.Println("Vera:", n, found) // Vera 2 true
	n, found = slices.BinarySearch(arr, "Bill")
	fmt.Println("Bill:", n, found) //Bill 1 false

	// 10. Delete
	arr = []string{"a", "b", "c", "d", "e"}
	arr = slices.Delete(arr, 1, 4)
	fmt.Println(arr) //[a e]

	// 11. Replace
	arr = []string{"Alice", "Bob", "Vera", "Zac"}
	arr = slices.Replace(arr, 1, 3, "Bill", "Billie", "Cat")
	fmt.Println(arr) //[Alice Bill Billie Cat Zac]

	// ********************************************************************************************

	// // 2. Sorting (Old Methods)

	// // sort int
	// highScores2 := make([]int, 0)
	// highScores2 = append(highScores2, 1, 5, 6, 4, 3, 5, 6)

	// sort.Ints(highScores2)
	// fmt.Println(highScores2)                     //[1 3 4 5 5 6 6]
	// fmt.Println(sort.IntsAreSorted(highScores2)) //true

	// // sort strings
	// highScores3 := make([]string, 0)

	// highScores3 = append(highScores3, "sai", "anand")

	// sort.Strings(highScores3)
	// fmt.Println(highScores3)
	// fmt.Println(sort.StringsAreSorted(highScores3)) //true

	// //Sort anykind

	// var names = []string{"sai", "anand", "kiki"}

	// sort.Slice(names, func(i, j int) bool {
	// 	return names[i] < names[j] //ascending
	// 	// return names[i] > names[j] //descending
	// })

	// fmt.Println(names) //[anand kiki sai]

}
