# WHAT I LEARNT ABOUT GO MAP

* I learnt that A map is an unordered and changeable collection that does not allow duplicates. And it is used to store data values in key:value pair.
* Each element in a map is a key:value pair.

## CREATION OF MAP USING var AND :=

Example

This example shows how to create maps in Go. Notice the order in the code and in the output
package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

* The map key can be of any data type for which the equality operator (==) is defined. These include:

 1. Booleans
 2. Numbers
 3. Strings
 4. Arrays
 5. Pointers
 6. Structs
 7. Interfaces (as long as the dynamic type supports equality)

* Invalid key types are:

 1. Slices
 2. Maps
 3.Functions

* These types are invalid because the equality operator (==) is not defined for them.
* The map value can be of any type. And you can access map elements by: value = map_name[key].

Example
package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Printf(a["brand"])
}
