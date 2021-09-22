package exercise7_8

import (
	"testing"
)

type ThreeColumns struct {
	A, B, C int
}

type StableTable struct {
	data []*ThreeColumns
}

func TestStableSortTracks(t *testing.T) {
	var table = StableTable{
		[]*ThreeColumns{{1, 2, 3}},
	}
	// want will show what we expect the order to be
	// ordering by column A then column B is it different if the ordering is not kept?
	// as we put tests together we can validate it.
}
