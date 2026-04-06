# WHAT I LEARNT FROM THE GO COMMENT

* I learnt that comment is a text that is ignored upon execution. And it helps in the readability and explaination of code. It's also used to prevent code execution when been text with alternative code. It may be a single-line or multi-line comments.

## Go Single-line Comments

* Single-line comments start with two forward slashes (//).Any text between // and the end of the line is ignored by the compiler (will not be executed).

Example 
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!") // This is a comment
}

## Go Multi-line Comments

* Multi-line comments start with /* and ends with */.Any text between /* and */ will be ignored by the compiler:

Example
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}
