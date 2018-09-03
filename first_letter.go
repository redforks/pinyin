package pinyin

import "sort"

type record struct {
	Order uint16
	// encoded first letter of pinyin, at most three letters for multi tones
	// each letters is 5bits, 'a' is 1, 'b' is 2, etc.
	// lower bits are first
	Tones uint16
}

func (r record) FirstLetters() []rune {
	return []rune{}
}

func FirstLetters(c rune) []rune {
	int16C := uint16(c)
	idx := sort.Search(len(records), func(idx int) bool {
		return records[idx].Order >= int16C
	})
	if idx == -1 || records[idx].Order != int16C {
		return []rune{}
	}

	return records[idx].FirstLetters()
}
