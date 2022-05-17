package ch7ex15

import (
	"bufio"
	"fmt"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ch7ex15 reads a single expression from the standard input,
// prompts the user to provide values for any variables,
// then evaluates the expression in the resulting environment.
func Ch7ex15() {
	expression := userInput("")
	if expression == "" {
		log.Fatalf("quit")
	}
	log.Printf("Evaluating expression: %q", expression)
	ex, err := expr.Parse(expression)
	if err != nil {
		log.Fatalf("Parse error: %s", err)
	}
	vars := make(map[expr.Var]bool)
	err = ex.Check(vars)
	if err != nil {
		log.Fatalf("Check error: %s", err)
	}
	env := make(expr.Env)
	for name := range vars {
		value, _ := strconv.ParseFloat(userInput(name.String()), 64)
		env[name] = value
	}
	fmt.Printf("result> %f\n", ex.Eval(env))
}

// TODO Use DCI model in userInput

func userInput(prompt string) (input string) {
	inputScanner := bufio.NewScanner(os.Stdin)
	for input == "" {
		fmt.Printf("%s? ", prompt)
		if !inputScanner.Scan() {
			if err := inputScanner.Err(); err != nil {
				log.Fatalf("Input error: %s", err)
			}
			break
		}
		input = strings.TrimSpace(inputScanner.Text())
	}
	return
}
