# Go Constants

* I learnt that if a variable should have a fixed value that cannot be changed, i can use the const keyword.The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.Example: const CONSTNAME type = value.

Here is an example of declaration of a constant in Go:

Example
package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}

## Constant Rules

 1. I learnt that Constant names follow the same naming rules as variables
 2. Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
 3. Constants can be declared both inside and outside of a function

### Constant Types

* And i learnt There are two types of constants:
 1. Typed constants
 2. Untyped constants

* I learnt that Typed Constants are declared with a defined type:

Example
package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}

* I learnt that untyped constants are declared without a type:

Example
package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}

Example
package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}

* I learnt that multiple constants can be grouped together into a block for readability:

Example
package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}

* I learnt that, when a constant is declared, it is not possible to change the value later:
