# WHAT I LEARNT IN GO SYNTAX

* I learnt that in Go, every program is start with a package declaration, importation of packages from a go libary, a white space, functions, statements and expressions.

## EXAMPLE
package main
import ("fmt")

func main() {
    fmt.Println("Hello Word!")
}

### EXPLAINATION
* Line 1; package main.
Here occurs the declaration of the package. in Go, every program is part of a package. in this case the program belongs to the main package.

* Line 2: import ("fmt").
Here comes the importation of packages. in this case, imported a package from a libary known as fmt. it is a shorthand for format.

* Line 3: a blank line.
it occurs for the readability. Go ignores it.

* Line 4: func main() {.
Here occurs the functions. in programming we use func, it is a shorthand used in replacement of function. the func main() indicated that the name of the function is main and dose not take in anything neither return anything that's why front of the main there is an empty brackets. 

* Line 5:fmt.Println("Hello Word!") }.
Here occurs the statements and executions. the fmt.Println("Hello Word!") is a function made available from the fmt package. and it's used to output or print text. the open and close curly brackets signifies execution.any thing that's inside the curly brackets will be executed. in the given example, it outputed "Hello Word!".

* I learnt that every executable code belongs to the main package.
