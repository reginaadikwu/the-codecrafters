# WHAT I LEARNT ABOUT SWITCH STATEMENT

* I learnt that the statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement. it si use to select one of many code blocks to be executed.

## Single-Case switch Syntax

Syntax
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}

* This is how it works:

 1. The expression is evaluated once
 2. The value of the switch expression is compared with the values of each case
 3. If there is a match, the associated block of code is executed
 4. The default keyword is optional. It specifies some code to run if there is no case match

 Example
package main
import ("fmt")

func main() {
  day := 8

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Not a weekday")
  }
}

## The Multi-case switch Statement

* I learnt it's  possible to have multiple values for each case in the switch statement.

Example
package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}
