package pinyin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/redforks/pinyin"
)

var _ = Describe("GetRuneFirstLetters", func() {
	It("GetRuneFirstLetters", func() {
		Ω(GetRuneFirstLetters('大')).Should(Equal([]rune{'d', 't'}))
		Ω(GetRuneFirstLetters('㐲')).Should(Equal([]rune{'d'}))
		Ω(GetRuneFirstLetters('〇')).Should(Equal([]rune{'l', 'y', 'x'}))

		Ω(GetRuneFirstLetters('a')).Should(Equal([]rune{'a'}))

		Ω(GetRuneFirstLetters('\n')).Should(Equal([]rune{}))
	})

	It("FirstLetters", func() {
		Ω(FirstLetters("a", ' ')).Should(Equal("a"))
		Ω(FirstLetters("a全\n", ' ')).Should(Equal("aq"))

		Ω(FirstLetters("a大\nb", ' ')).Should(Equal("adb atb"))
		Ω(FirstLetters("a大大", ' ')).Should(Equal("add adt atd att"))
		Ω(FirstLetters("大a大", ' ')).Should(Equal("dad dat tad tat"))
		Ω(FirstLetters("大虾㚘", ' ')).Should(Equal("dxb dxh dxf dhb dhh dhf txb txh txf thb thh thf"))
	})
})
