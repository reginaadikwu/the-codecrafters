# WHAT I LEARNT ABOUT SLICES

* I learnt that slices are similar to arrays, but they are more powerful and flexible. 
* Like arrays, slices are also used to store multiple values of the same type in a single variable.
* However, unlike arrays, the length of a slice can grow and shrink as you see fit.

* And i learnt that in Go, there are several ways of creating a slice:
 1. Using the []datatype{values} format

* This example shows how to create slices using the []datatype{values} format:
package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}

 2. Create a slice from an array

* This example shows how to create a slice from an array:
package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}
 3. Using the make() function

This example shows how to create slices using the make() function:
package main
import ("fmt")

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}

