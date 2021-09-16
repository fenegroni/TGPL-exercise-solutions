package exercise7_6

import (
	tempconv "TGPL-exercise-solutions/chapter7/exercise7.6/temconv"
	"flag"
	"fmt"
)

func print() {
	flag.Parse()
	fmt.Print(tempconv.AbsoluteZeroC)
}
