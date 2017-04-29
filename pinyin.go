package pinyin

import (
	"bytes"
	"regexp"
	"strings"
)

// VERSION defines the running build id.
var VERSION = "0.20.0"
var buildTime = "2017-04-30"

// Meta
const (
	Authors = `
mozillazg, 闲耘
suntong
`
	License   = "MIT"
	Copyright = `
Copyright (c) 2016 mozillazg, 闲耘
Copyright (c) 2017 suntong
`
)

// == 拼音风格
const (
	Normal = iota // 普通风格，不带声调（默认风格）。如： zhong guo
)

// -- 声调风格 Tone
const (
	_     = iota
	Tone1 // 声调风格1，即拼音声调在各个拼音之后，用数字 [1-4] 进行表示。如： zhong1 guo2
	Tone2 // 声调风格2，即拼音声调在各个韵母之后，用数字 [1-4] 进行表示。如： zho1ng guo2
	Tone3 // 声调风格3，拼音声调在韵母上。如： zhōng guó
)

// -- 部分返回 Truncate
const (
	_           = iota
	FirstLetter     // 1: 首字母风格，只返回拼音的首字母部分。如： z g
	Initials        // 2: 声母风格，只返回各个拼音的声母部分。如： zh g
	Finals      = 9 // 9: 韵母风格，只返回各个拼音的韵母部分。如： ong uo
)

// 声母表
var initialArray = strings.Split(
	"b,p,m,f,d,t,n,l,g,k,h,j,q,x,r,zh,ch,sh,z,c,s",
	",",
)

// 所有带声调的字符
var rePhoneticSymbolSource = func(m map[string]string) string {
	s := ""
	for k := range m {
		s = s + k
	}
	return s
}(phoneticSymbol)

// 匹配带声调字符的正则表达式
var rePhoneticSymbol = regexp.MustCompile("[" + rePhoneticSymbolSource + "]")

// 匹配使用数字标识声调的字符的正则表达式
var reTone2 = regexp.MustCompile("([aeoiuvnm])([1-4])$")

// 匹配 Tone2 中标识韵母声调的正则表达式
var reTone1 = regexp.MustCompile("^([a-z]+)([1-4])([a-z]*)$")

// Style 配置拼音风格 (声调风格 + 部分返回)
type Style struct {
	Tone     int // 拼音风格（默认： Normal)
	Truncate int // 部分返回
}

// Pinyin with 配置信息
type Pinyin struct {
	Style
	Polyphone bool   // 是否启用多音字模式（默认：禁用）
	Separator string // 使用的分隔符（默认：" ")
}

// Separator 默认配置：所用的分隔符
var Separator = " "

var finalExceptionsMap = map[string]string{
	"ū": "ǖ",
	"ú": "ǘ",
	"ǔ": "ǚ",
	"ù": "ǜ",
}
var reFinalExceptions = regexp.MustCompile("^(j|q|x)(ū|ú|ǔ|ù)$")
var reFinal2Exceptions = regexp.MustCompile("^(j|q|x)u(\\d?)$")

// NewPinyin 返回包含默认配置的 `Pinyin`
func NewPinyin() Pinyin {
	return Pinyin{Separator: Separator}
}

func (a *Pinyin) SetStyle(s Style) {
	a.Tone = s.Tone
	a.Truncate = s.Truncate
}

// 获取单个拼音中的声母
func initial(p string) string {
	s := ""
	for _, v := range initialArray {
		if strings.HasPrefix(p, v) {
			s = v
			break
		}
	}
	return s
}

// 获取单个拼音中的韵母
func final(p string) string {
	n := initial(p)
	if n == "" {
		return handleYW(p)
	}

	// 特例 j/q/x
	matches := reFinalExceptions.FindStringSubmatch(p)
	// jū -> jǖ
	if len(matches) == 3 && matches[1] != "" && matches[2] != "" {
		v, _ := finalExceptionsMap[matches[2]]
		return v
	}
	// ju -> jv, ju1 -> jv1
	p = reFinal2Exceptions.ReplaceAllString(p, "${1}v$2")
	return strings.Join(strings.SplitN(p, n, 2), "")
}

// 处理 y, w
func handleYW(p string) string {
	// 特例 y/w
	if strings.HasPrefix(p, "yu") {
		p = "v" + p[2:] // yu -> v
	} else if strings.HasPrefix(p, "yi") {
		p = p[1:] // yi -> i
	} else if strings.HasPrefix(p, "y") {
		p = "i" + p[1:] // y -> i
	} else if strings.HasPrefix(p, "wu") {
		p = p[1:] // wu -> u
	} else if strings.HasPrefix(p, "w") {
		p = "u" + p[1:] // w -> u
	}
	return p
}

func (a Pinyin) toFixed(p string) string {
	if a.Truncate == Initials {
		return initial(p)
	}
	origP := p

	// 替换拼音中的带声调字符
	py := rePhoneticSymbol.ReplaceAllStringFunc(p, func(m string) string {
		symbol, _ := phoneticSymbol[m]
		switch a.Tone {
		// 不包含声调
		case Normal:
			// 去掉声调: a1 -> a
			m = reTone2.ReplaceAllString(symbol, "$1")
		case Tone2, Tone1:
			// 返回使用数字标识声调的字符
			m = symbol
		default:
			// 声调在头上
		}
		return m
	})

	switch a.Tone {
	// 将声调移动到最后
	case Tone1:
		py = reTone1.ReplaceAllString(py, "$1$3$2")
	}
	switch a.Truncate {
	// 首字母
	case FirstLetter:
		py = py[:1]
	// 韵母
	case Finals:
		// 转换为 []rune unicode 编码用于获取第一个拼音字符
		// 因为 string 是 utf-8 编码不方便获取第一个拼音字符
		rs := []rune(origP)
		switch string(rs[0]) {
		// 因为鼻音没有声母所以不需要去掉声母部分
		case "ḿ", "ń", "ň", "ǹ":
		default:
			py = final(py)
		}
	}
	return py
}

// Convert 汉字转拼音，支持多音字模式.
// If enabled Polyphone, then separate the returns with '/'.
// E.g., for input like "我的银行不行", the output is
// wo de yin hang/xing bu hang/xing.
func (a Pinyin) Convert(s string) string {
	pys := bytes.NewBufferString("")
	for _, r := range s {
		if r <= '~' {
			pys.WriteString(string(r))
			continue
		}
		value, ok := PinyinDict[int(r)]
		if !ok {
			pys.WriteString(string(r))
			continue
		}
		firstComma := strings.Index(value, ",")
		if !a.Polyphone && firstComma > 0 {
			value = value[:firstComma]
		}
		// 多音字模式 (Polyphone), output likes "hang/xing"
		if a.Polyphone && firstComma > 0 {
			value = strings.Replace(value, ",", "/", -1)
		}
		py := a.toFixed(value)
		pys.WriteString(py + a.Separator)
	}
	return pys.String()
}
