# WHAT I LEARNT ABOUT VARIABLES

*I learnt that variables are empty containers that stores data values. In Go, there differnt types of variables, these are:

1. int; this stores whole numbers, such as 123(positive numbers) or -123(negetive numbers).

2. float32; this stores floating point numbers, with decimals, such as 19.99(positive decimal numbers) or -19.99(nagetive decimal numbers).

3. string; this stores text, such as "Hello World". String values are surrounded by double quotes.

4. bool; this  stores values with two states: true or false.

## Declaration Of Variables

* I learnt that in Go, there are two ways to declare a variable:
1. With the var keyword:
Use the var keyword, followed by variable name and type:
Example: var variablename type = value

2. With the := sign:
Use the := sign, followed by the variable value: 
Example: variablename := value

### Variable Declaration With Initial Value
* I learnt that if the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:
Example
package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}

#### Variable Declaration Without Initial Value
* I learnt that in Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:
Example
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
In this example there are 3 variables:

    a
    b
    c

* These variables are declared but they have not been assigned initial values.By running the code, we can see that they already have the default values of their respective types:

    a is ""
    b is 0
    c is false

##### Value Assignment After Declaration

* I learnt that it is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.
Example
package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}

###### Difference Between var and :=

* There are some small differences between the var var :=:
var
1. Can be used inside and outside of functions. 
2. Variable declaration and value assignment can be done separately. WHILE

:=
1. Can only be used inside functions.
2. Variable declaration and value assignment cannot be done separately (must be done in the same line)

This example shows declaring variables outside of a function, with the var keyword:
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

# Go Multiple Variable Declaration

* I learnt that in Go, it is possible to declare multiple variables on the same line.

This example shows how to declare multiple variables on the same line:
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

* I learnt that if the type keyword is not specified, you can declare different types of variables on the same line:

Example
package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

## Go Variable Declaration in a Block

* I learnt that multiple variable declarations can also be grouped together into a block for greater readability:

Example
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

# Go Variable Naming Rules

* I learnt that a variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

* I also learnt the Go variable naming rules:

 1.  A variable name must start with a letter or an underscore character (_)
 2.  A variable name cannot start with a digit
 3.  A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
 4.  Variable names are case-sensitive (age, Age and AGE are three different variables)
 5.  There is no limit on the length of the variable name
 6.  A variable name cannot contain spaces
 7.  The variable name cannot be any Go keywords

## Multi-Word Variable Names

* I learnt that variable names with more than one word can be difficult to read.

* And also, There are several techniques that can be used to make them more readable:

# Camel Case
I learnt that each word, except the first, starts with a capital letter:
myVariableName = "goIsFun"

# Pascal Case
I learnt that each word starts with a capital letter:
MyVariableName = "JoyIsCute"

# Snake Case
I learnt that each word is separated by an underscore character:
my_variable_name = "he_hate_her".
