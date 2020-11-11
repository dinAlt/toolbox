package stringlist

import (
	"fmt"
	"strings"
)

const defaultIncrementCount = 10
const defaultCapacity = 10

type Strings interface {
	Rows() []string
}

// StringList is set of string methods for slice
type StringList struct {
	rws     []string
	lineSep string
}

func New(capacity int) *StringList {
	if capacity == 0 {
		capacity = defaultCapacity
	}
	return &StringList{lineSep: "\n", rws: make([]string, 0, capacity)}
}

func (sl *StringList) Add(row string) {
	if cap(sl.rws) == len(sl.rws) {
		sl.rws = append(make([]string, 0, len(sl.rws)+defaultIncrementCount), sl.rws...)
	}
	sl.rws = append(sl.rws, row)
}

func (sl *StringList) FmtAdd(row string, a ...interface{}) {
	sl.Add(fmt.Sprintf(row, a...))
}

func (sl *StringList) AddRows(rows []string) {
	if cap(sl.rws)-len(sl.rws) < len(rows) {
		sl.rws = append(make([]string, 0, len(sl.rws)+len(rows)+defaultIncrementCount), sl.rws...)
	}
	sl.rws = append(sl.rws, rows...)
}

func (sl *StringList) AddRange(rn ...string) {
	sl.AddRows(rn)
}

func (sl *StringList) AddFrom(list Strings) {
	sl.AddRows(list.Rows())
}

func (sl *StringList) RowAt(index int) (string, error) {
	if (index > len(sl.rws)-1) || (index < 0) {
		return "", fmt.Errorf("StringList.RowAt: list index ot of bounds len: %d, idx: %d", len(sl.rws), index)
	}
	return sl.rws[index], nil
}

func (sl *StringList) Rows() []string {
	return append(make([]string, 0), sl.rws...)
}

func (sl *StringList) Text() string {
	return strings.Join(sl.rws, sl.lineSep)
}

func (sl *StringList) Join(sep string) string {
	return strings.Join(sl.rws, sep)
}

func (sl *StringList) SetLineSep(sep string) {
	sl.lineSep = sep
}

func (sl *StringList) Len() int {
	return len(sl.rws)
}

func (sl *StringList) ReplaceAt(index int, row string) error {
	if (index > len(sl.rws)-1) || (index < 0) {
		return fmt.Errorf("StringList.ReplaceAt: list index ot of bounds len: %d, idx: %d", len(sl.rws), index)
	}
	sl.rws[index] = row
	return nil
}
