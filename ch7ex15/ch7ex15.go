package ch7ex15

import (
	"fmt"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"
	"log"
)

// Ch7ex15 reads a single expression from the standard input,
// prompts the user to provide values for any variables,
// then evaluates the expression in the resulting environment.
func Ch7ex15() {
	type varMap map[expr.Var]bool
	var inEx string
	fmt.Print("?> ")
	// TODO handle error
	fmt.Scanf("%s", &inEx)
	ex, err := expr.Parse(inEx)
	if err != nil {
		log.Fatalf("error parsing the expression: %s", err)
	}
	vars := make(varMap)
	err = ex.Check(vars)
	if err != nil {
		log.Fatalf("error checking the expression: %s", err)
	}
	// TODO ask for values for each variable
}
