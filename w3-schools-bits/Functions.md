# WHAT I LEARNT ABOUT FUNCTION

* I learnt a function is a block of statements that can be used repeatedly in a program. It will not be executed automatically when a page loads, It will be executed by a call of it.

* I learnt that Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

* In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():

package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}

## Naming Rules for Go Functions

 1. A function name must start with a letter
 2. A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
 3. Function names are case-sensitive
 4. A function name cannot contain spaces
 5. If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used
 
 ### Parameters and Arguments

* I learnt that information can be passed to functions as a parameter. And parameters act as variables inside the function.

* And parameters and it types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma.
Example
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}

#### Return Values

*I learnt that if you want a function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function.
Example
Here, myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int):
package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}

##### Recursion Functions

* I learnt that Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use the x variable as the data, which increments with 1 (x + 1) every time we recurse. The recursion ends when the x variable equals to 11 (x == 11). 
Example
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}

