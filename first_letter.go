package pinyin

import (
	"sort"
	"strings"
	"unicode"
)

type record struct {
	CodePoint uint16
	// encoded first letter of pinyin, at most three letters for multi tones
	// each letters is 5bits, 'a' is 1, 'b' is 2, etc.
	// lower bits are first
	Tones uint16
}

func (r record) FirstLetters() []rune {
	var rv = make([]rune, 0, 3)
	for i := (uint16)(0); i < 3; i++ {
		bits := (r.Tones >> (5 * i)) & 0x1f
		if bits == 0 {
			break
		}
		bits--
		rv = append(rv, (rune)(bits+'a'))
	}
	return rv
}

// GetRuneFirstLetters get char pinyin first letter, multiple result if it has multi tones,
// start with most common tone. Return itself if it is not chinese.
// Return empty slice if not printable char.
func GetRuneFirstLetters(c rune) []rune {
	if !unicode.IsPrint(c) {
		return []rune{}
	}

	int16C := uint16(c)
	idx := sort.Search(len(records), func(idx int) bool {
		return records[idx].CodePoint >= int16C
	})
	if idx == -1 || records[idx].CodePoint != int16C {
		return []rune{c}
	}

	return records[idx].FirstLetters()
}

// FirstLetters returns first letters of a string,
// If char not chinese, return char itself.
// If char has multiple tone, returns all combinations, separate by sep argument.
func FirstLetters(s string, sep string) string {
	runeSet := [][]rune{}
	for _, c := range s {
		rs := GetRuneFirstLetters(rune(c))
		if len(rs) != 0 {
			runeSet = append(runeSet, rs)
		}
	}

	combinations := 1
	dims := make([]int, len(runeSet))
	dimIdxs := make([]int, len(dims))
	for i := 0; i < len(dims); i++ {
		dims[i] = len(runeSet[i])
		combinations *= len(runeSet[i])
	}

	rvs := []string{}

	appendResult := func() {
		rec := make([]rune, len(dims))
		for i := 0; i < len(dims); i++ {
			rec[i] = runeSet[i][dimIdxs[i]]
		}

		s = string(rec)
		if !contains(rvs, s) {
			rvs = append(rvs, s)
		}
	}

	var addAll func(int)
	addAll = func(startIdx int) {
		idx := startIdx
		for idx = startIdx; idx < len(dims); idx++ {
			if dims[idx] != 1 {
				break
			}
		}

		appendResult()
		if idx < len(dims) {
			for i := 0; i < dims[idx]; i++ {
				for j := idx + 1; j < len(dims); j++ {
					dimIdxs[j] = 0
				}

				addAll(idx + 1)
				dimIdxs[idx]++
			}
		}
	}

	addAll(0)

	return strings.Join(rvs, sep)
}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
