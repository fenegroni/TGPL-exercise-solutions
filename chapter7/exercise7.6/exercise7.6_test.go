package exercise7_6

import (
	tempconv "TGPL-exercise-solutions/chapter7/exercise7.6/temconv"
	"flag"
	"fmt"
	"testing"
)

func TestSetCelsius(t *testing.T) {
	flag.Parse()
	fmt.Println(tempconv.AbsoluteZeroC)
}
