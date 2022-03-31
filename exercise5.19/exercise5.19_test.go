package exercise5_19

import "testing"

func TestNoReturnStatement(t *testing.T) {
	arg := 1978
	defer func() {
		if arg != recover() {
			t.Fatalf("Function did not return argument.")
		}
	}()
	NoReturnStatement(arg)
}
