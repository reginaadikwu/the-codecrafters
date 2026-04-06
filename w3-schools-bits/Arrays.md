# WHAT I LEARNT ABOUT GO ARRAYS

* I learnt that arrays are used to store multiple values of the same type
in a single variable, instead of declaring separate variables for each value.

* And i learnt that in Go, there are two ways of declaring an array. These  are:
1. With the var keyword: examples
var array_name = [length]datatype{values} // here length is defined
OR
var array_name = [...]datatype{values} // here length is infe

2. With the := sign: examples
array_name := [length]datatype{values} // here length is defined
OR
array_name := [...]datatype{values} // here length is inferred 

EXAMPLE 
This example declares two arrays (arr1 and arr2) with defined lengths:
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}

WHILE

This example declares two arrays (arr1 and arr2) with inferred lengths:
package main
import ("fmt")

func main() {
  var arr1 = [...]int{1,2,3}
  arr2 := [...]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
