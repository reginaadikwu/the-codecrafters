# WHAT I LEARNT ABOUT OPERATORS

* I learnt that operators are used to perform operations on variable and values. The + operator add two values together, like the ezample given below.
package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a)
}

* Although the + operator is often used to add together two values, it can also be used to add together a variable and a value, or a variable and another variable.

Example
package main
import ("fmt")

func main() {
  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
} 

* I learnt that Go divides the operators into the following groups:
 1.  Arithmetic operators: which are used to perform common mathematical operations.
 2.  Assignment operators: which are used to assign values to variables.
 3.  Comparison operators: which are used to compare two values.
 4.  Logical operators: which are used to determine the logic between variables or values. 
 5.  Bitwise operators: which are used no (binary) numbers. 


