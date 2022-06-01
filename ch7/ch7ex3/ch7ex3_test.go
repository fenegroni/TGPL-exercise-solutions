package ch7ex3

import "testing"

func TestTree_String(t *testing.T) {
	var z *tree
	want := "[ ]"
	if got := z.String(); got != want {
		t.Logf("z = %s, want %s", got, want)
	}
	for _, v := range []int{10, 1, 12, 6, 2} {
		z = add(z, v)
	}
	want = "[ 1 2 6 10 12 ]"
	if got := z.String(); got != want {
		t.Logf("z = %s, want %s", got, want)
	}
}
