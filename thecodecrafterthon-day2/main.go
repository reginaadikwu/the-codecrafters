// CodeCrafters — Hackathon 002
// Squad: [The Channels]
/*
All Members Names:[
1.  Blessing Ogbaka
2.	Okopi Ebo Joy
3.  Daniel Akpa
4.  Antony Agbo
5.	Adikwu Regina
6.  Raymond Nicholas
7.  Jonathan Ahubi
8.	Daniel Okoh
9.  Moses Ochife
]
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

back:
	fmt.Println("════════════════════════════════════════════════")
	fmt.Println(" SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("All systems nominal. Type 'help' to begin.")
	fmt.Println("════════════════════════════════════════════════")

	fmt.Print("C&C> ")
	var history []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	cmd := strings.Fields(input)

	inputed := scanner.Text()

	switch strings.ToLower(inputed) {

	case "help":
		fmt.Println("CALC HELP")
		fmt.Println("calc add <a> <b>    -> addition")
		fmt.Println("calc sub <a> <b>    -> substraction")
		fmt.Println("calc mul <a> <b>    -> multiplication")
		fmt.Println("calc div <a> <b>    -> division")
		fmt.Println("calc mod <a> <b>    -> ramainder")
		fmt.Println("calc pow <a> <b>    -> a to the power of b")
		fmt.Println("calc last           -> show the last result")
		fmt.Println("calc history        -> show last 5 calc results\n\n")

		fmt.Println("BASE HELP")
		fmt.Println("C&C> base dec 255✦ \n Binary : 11111111 \n ✦ Hex    : FF")
		fmt.Println(" C&C> base hex 1F \n✦ Decimal: 31\n")
		fmt.Println("C&C> base bin 1010 \n✦ Decimal: 10\n")

		fmt.Println("STRING TRANSFORMER HELP")
		fmt.Println(" str upper  <text>    → ALL UPPERCASE")
		fmt.Println("str lower  <text>    → all lowercase")
		fmt.Println("str cap    <text>    → Every Word Capitalised")
		fmt.Println("str title  <text>    → Title Case With Small Words")
		fmt.Println("str snake  <text>    → convert_to_snake_case")
		fmt.Println("str reverse <text>   → esreveR hcaE droW")
		goto back
	case "quit":
		fmt.Println("THANK YOU AND GOOD BYE")
		return
	}

	// if len(cmd) >=5 {
	// 	fmt.Println("Invalid number of arguments")
	// 	return
	// }

	if cmd[0] != "calc <a> <b>" && cmd[1] != "upper" && cmd[1] != "lower" && cmd[1] != "cap" && cmd[1] != "title" && cmd[1] != "snake" && cmd[1] != "reverse" && cmd[1] != "add" && cmd[1] != "sub" && cmd[1] != "mul" && cmd[1] != "div" && cmd[1] != "mod" && cmd[1] != "pow" && cmd[1] != "help" && cmd[1] != "quit" {
		fmt.Println("Invalid command")
		goto back
	}

	switch strings.ToLower(cmd[1]) {

	case "upper":

		new := cmd[2:]

		result := strings.Join(new, " ")

		fmt.Println(strings.ToUpper(result))
		history = append(history, "upper "+strings.ToUpper(result))

		goto back

	case "lower":
		new := cmd[2:]

		result := strings.Join(new, " ")

		fmt.Println(strings.ToLower(result))
		history = append(history, "lower "+strings.ToLower(result))

		goto back

	case "cap":
		new2 := cmd[2:]

		result := strings.Join(new2, " ")

		fmt.Println(strings.Title(result))
		history = append(history, "cap "+strings.Title(result))

		goto back
	case "reverse":
		reverse := cmd[2:]
		var result []string
		for i := len(reverse) - 1; i >= 0; i-- {
			result = append(result, reverse[i])
		}
		fmt.Println(strings.Join(result, ""))

		goto back

	case "add":
		if len(cmd) != 4 {
			fmt.Println("ERORR!! ENTER HELP FOR SUPPORT")
			goto back
		}

		add1, _ := strconv.ParseInt(cmd[2], 10, 64)
		add2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result: ", add1+add2)

		goto back

	case "sub":
		if len(cmd) != 4 {
			fmt.Println("ERORR!! ENTER HELP FOR SUPPORT")
			goto back
		}
		sub1, _ := strconv.ParseInt(cmd[2], 10, 64)
		sub2, _ := strconv.ParseInt(cmd[3], 10, 64)
		fmt.Println("✦ Result: ", sub1-sub2)
		goto back

	case "mul":
		if len(cmd) != 4 {
			fmt.Println("ERORR!! ENTER HELP FOR SUPPORT")
			goto back
		}
		mul1, _ := strconv.ParseInt(cmd[2], 10, 64)
		mul2, _ := strconv.ParseInt(cmd[3], 10, 64)
		fmt.Println("✦ Result: ", mul1*mul2)

	case "div":
		if len(cmd) != 4 {
			fmt.Println("ERORR!! ENTER HELP FOR SUPPORT")
			goto back
		}
		div1, _ := strconv.ParseInt(cmd[2], 10, 64)
		div2, _ := strconv.ParseInt(cmd[3], 10, 64)
		if div2 == 0 {
			fmt.Println("can't be divided by ZERO")
			return
		}
		fmt.Println("✦ Result: ", float64(div1)/float64(div2))
		goto back

	case "mod":
		if len(cmd) != 4 {
			fmt.Println("ERORR!! ENTER HELP FOR SUPPORT")
			goto back
		}
		mod1, _ := strconv.ParseInt(cmd[2], 10, 64)
		mod2, _ := strconv.ParseInt(cmd[3], 10, 64)
		if mod2 == 0 {
			fmt.Println("error")
			return
		}
		fmt.Println("✦ Result: ", mod1%mod2)
		goto back

	case "pow":
		if len(cmd) != 4 {
			fmt.Println("ERORR!! ENTER HELP FOR SUPPORT")
			goto back
		}
		pow1, _ := strconv.ParseInt(cmd[2], 10, 64)
		pow2, _ := strconv.ParseInt(cmd[3], 10, 64)
		fmt.Println("✦ Result: ", math.Pow(float64(pow1), float64(pow2)))
		goto back
	case "history":
		if len(history) <= 5 {
			fmt.Println(history)
			goto back
		}
		fmt.Println(history[len(history)-5], "\n", history[len(history)-4], "\n", history[len(history)-3], "\n", history[len(history)-2], "\n", history[len(history)-1], "\n")

	}

}
