package main

import "fmt"

//func main() {
//	// Todo: Fix this program. Uncomment the lines below, delete or comment out "others"
//	var students = 50
//	var assistants = 15
//	var professors = 5
//	//var others = 1
//
//	println(students + assistants)
//	_ = professors //=>setting it to undefined, drop the variable, used for debugging
//}

//func main() {
//	// Todo: Can you move line 6 into the loop header? What's the difference?
//	var i = 0
//
//	for i = 0; i < 10; i++ {
//		println(i)
//	}
//}

func moin() {
	println("moin")
}

//func main() {
//	var i int = 0
//	for i = 1; i < 3; i++ {
//		moin()
//	}
//}

//

// Example: appending items to a slice
func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
