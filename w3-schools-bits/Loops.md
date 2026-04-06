# WHAT I LAERNT ABOUT FOR LOOPS

* I learnt that Loops are handy if you want to run the same code over and over again, each time with a different value. Its execution is called an iteration. And it is the only loop avilable in Go.

This example will print the numbers from 0 to 4:  
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}

Explaination of example
 * i:=0; - Initialize the loop counter (i), and set the start value to 0
 * i < 5; - Continue the loop as long as i is less than 5
 * i++ - Increase the loop counter value by 1 for each iteration

 ## The continue Statement

 * I learnt Its used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

 This example skips the value of 3:
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}

### The break Statement

* I learnt iteration is used to break/terminate the loop execution.

This example breaks out of the loop when i is equal to 3:
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}

#### Nested Loops

* I learnt it is possible to place a loop inside another loop.
Here, the "inner loop" will be executed one time for each iteration of the "outer loop":
Example
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}

##### The Range Keyword

* I learnt the range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.
The range keyword is used like this:for index, value := range array|slice|map {
   // code to be executed for each iteration
}
Example
This example uses range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}
