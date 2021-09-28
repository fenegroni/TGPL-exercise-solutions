package exercise7_8

import (
	"fmt"
	"strings"
)

type TableRow struct {
	col int
	pos int
}

type Table []*TableRow

func (t Table) Len() int             { return len(t) }
func (t Table) Less(i, j int) bool   { return t[i].col < t[j].col }
func (t Table) Swap(i, j int)        { t[i], t[j] = t[j], t[i] }
func (t Table) At(i int) interface{} { return t[i] }

func (t Table) String() string {
	var buf strings.Builder
	for i := range t {
		buf.WriteString(fmt.Sprintf("{col:%d,pos:%d} ", t[i].col, t[i].pos))
	}
	return buf.String()
}
