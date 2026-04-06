package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("-------Hello Dear, Welcome. How may i be of help?-------")

	for {

		var input string
		fmt.Println("Type quit to exit")
		fmt.Println("---OR---")
		fmt.Println("Type HELP to enter menu")
		fmt.Scanln(&input)

		if input == "quit" {
			fmt.Println("Goodbye")
			break
		} else if input == "HELP" {
			fmt.Println("for Addition: Numbers, Operations, Numbers(1 + 1)")
			fmt.Println("for Subtraction: Numbers, Operations, Numbers(1 - 1)")
			fmt.Println("for Multiplication: Numbers, Operations, Numbers(1 * 1)")
			fmt.Println("for Division: Numbers, Operations, Numbers(1 / 1)")
			fmt.Println("Move On")

			goto Commence
		} else {
			fmt.Println("HangUp, Not Recorgnised.")
			break
		}

	Commence:

		var dex string
		fmt.Println("Enter First Number")
		fmt.Scanln(&dex)
		index1, err := strconv.Atoi(dex)
		if err != nil {
			fmt.Println("Only Number's Are Allowed!")
			fmt.Println("START OVER.")
			goto Commence
		}

		var operator string
		fmt.Println("enter operator")
		fmt.Scanln(&operator)
		if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			fmt.Println("Not An Operator")
			fmt.Println("START OVER")
			goto Commence
		}

		var Index string
		fmt.Println("Enter Second Number")
		fmt.Scanln(&Index)
		index2, err := strconv.Atoi(Index)
		if err != nil {
			fmt.Println("Only Number's Are Allowed!")
			fmt.Println("START OVER")
			goto Commence
		}

		if operator == "+" {
			fmt.Println("result:", index1+index2)
			goto Commence
		}

		if operator == "-" {
			fmt.Println("result:", index1-index2)
			goto Commence
		}

		if operator == "*" {
			fmt.Println("result:", index1*index2)
			goto Commence
		}

		if operator == "/" {
			if index1 == 0 || index2 == 0 {
				fmt.Println("Syntax Error")
			} else {
				fmt.Println("result:", index1/index2)
				goto Commence
			}

		}

	}

}
