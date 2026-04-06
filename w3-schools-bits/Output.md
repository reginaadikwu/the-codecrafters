# WHAT I LEARTN FROM GO OUTPUT FUNCTION

* I learnt that Go has three functions to output text And these are:

 1. Print()
 2. Println()
 3. Printf()

## The Print() Function

I learnt that the Print() function prints its arguments with their default format.

Example
Print the values of i and j:
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}

* And if you want to print the arguments in new lines, you need to use \n.

Example
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}

* And it is possible to only use one Print() for printing multiple variables.

Example
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n",j)
}

If you want to add a space between string arguments, you need to use " ".

Example
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, " ", j)
}

### The Println() Function

* I learnt that the Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end.

Example
package main
import ("fmt")

func main() {
  var a,b string = "Hello","World"

  fmt.Println(i,j)
}

#### The Printf() Function

* I learnt that the Printf() function first formats its argument based on the given formatting verb and then prints them.

Here two formatting verbs are used:
 1. %v is used to print the value of the arguments
 2. %T is used to print the type of the arguments

Example
package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}

# THE Go Formatting Verbs  For Printf().

* I learnt that Go offers several formatting verbs that can be used with the Printf() function.

1. # General Formatting Verbs
Here are the following verbs that can be used with all data types:

VERB    DESCRIPTION
%v      Prints the value in the default format
%#v     Prints the value in Go-syntax format
%T      Prints the type of the value
%%      Prints the % sign

Example
package main
import ("fmt")

func main() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)
}

