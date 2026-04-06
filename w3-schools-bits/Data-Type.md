# WHAT I LEARNT FROM DATA TYPE

* I learnt that Data type is an important concept in programming. Data type specifies the size and type of variable values.

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

Go has three basic data types:

   1. bool: represents a boolean value and is either true or false
   2. Numeric: represents integer types, floating point values, and complex types
   3. string: represents a string value

This example shows some of the different data types in Go:
package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)

  * And i also learnt that there are four data type, which are:
    1. boolean data type: it is declared with the bool keyword and can only take the values of true or false. with a default value is false.
    2. Integer data type: it is used to store a whole number without decimals, like 35, -50, or 1345000. with the categories of signed and unsigned integer.
      - The signed integers: this can store both positive and negative values.
      - The unsigned integers - this can only store non-negative values
    3. Float data type: this are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987. It has two keywords this are:
      - The float32 Keyword
      - The float64 Keyword
    4. The string data type: this is used to store a sequence of characters (text). And it values must be surrounded by double quotes.



}
