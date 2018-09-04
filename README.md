# Pinyin Library for Go!

## GetRuneFirstLetters

```go
func GetRuneFirstLetters(c rune) []rune
```

获取单个字符的拼音首字母:

1. 如果该字符为非打印字符, 则返回空 slice
1. 如果该字符非汉字, 返回字符本身, 如: `GetRuneFirstLetters('a') // []rune{'a'}`
1. 如果该字符为汉字, 返回拼音首字, 如: `GetRuneFirstLetters('人') // []rune{'r'}`
1. 如果该字符为汉字多音字, 则返回每个音的拼音首字, 常用音在前, 如: `GetRuneFirstLetters('大') // []ruse{'d', 't'}`

## FirstLetters

```go
func FirstLetters(s string, sep string) string
```

返回字符串的拼音首字母:

忽略不可打印字符, 非汉字返回字符自身:

```go
FirstLetters("a全\n", " ") // aq
```

如果包含多音字, 生成多音字的所以组合, 用`sep`字符隔开:

```go
FirstLetters("大虾㚘", " ") // dxb dxh dxf dhb dhh dhf txb txh txf thb thh thf
```

## 限制

多音字只保留最常用的三个读音, 实际上超过三个读音的多音字很少, 只有 10 几个, 且都是生僻音.

只包括 unicode 编码在 0xffff 之前的汉字, 基本都是生僻字, 不影响正常使用.
