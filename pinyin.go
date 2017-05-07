////////////////////////////////////////////////////////////////////////////
// Porgram: pinyin
// Purpose: pinyin conversion Go library
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Copyright (c) 2016 mozillazg, 闲耘
// 	        https://github.com/mozillazg/go-pinyin/
////////////////////////////////////////////////////////////////////////////

package pinyin

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-shaper/shaper"
)

// VERSION defines the running build id.
var (
	VERSION   = "0.20.0"
	buildTime = "2017-05-03"
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
	FirstLetter   = iota + 1 // 1: 首字母风格，只返回拼音的首字母部分。如： z g
	Initials                 // 2: 声母风格，只返回各个拼音的声母部分。如： zh g
	ZeroConsonant = iota + 6 // 8: 支持零声母功能, wén -> -> uén
	Finals                   // 9: 韵母风格，只返回各个拼音的韵母部分。如： ong uo, wén -> w -> én
	Both          = 11       // 11: 双显风格，返回 汉字 + 拼音
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

// Style 配置拼音风格 (声调风格 + 部分返回)
type Style struct {
	tone     int // 拼音风格（默认： Normal)
	truncate int // 部分返回
}

// Pinyin with 配置信息
type Pinyin struct {
	Style
	Separator   string // 使用的分隔符（默认：" ")
	polyphone   bool   // 是否启用多音字模式（默认：禁用）
	capitalized bool   // 首字母大写

	shaper *Shaper
}

var finalExceptionsMap = map[string]string{
	"ū": "ǖ",
	"ú": "ǘ",
	"ǔ": "ǚ",
	"ù": "ǜ",
}

// NewPinyin 返回包含默认配置的 `Pinyin`
func NewPinyin(tone, truncate int, separator string, _polyphone, _capitalized bool) Pinyin {
	a := Pinyin{Style: Style{tone, truncate},
		Separator:   separator,
		polyphone:   _polyphone,
		capitalized: _capitalized,
	}
	if a.truncate != ZeroConsonant {
		// 简明整齐的处理声母韵母 ref mozillazg/go-pinyin/issues/18
		// both y and w are considered 声母, add them back
		initialArray = append(initialArray, "y", "w")
	}
	a.shaper = NewShaper()
	if a.truncate != Normal {
		a.shaper.ApplyTruncate(a)
	}
	if a.tone != Tone3 {
		a.shaper.ApplyToneShaping(a)
	}
	if a.capitalized {
		a.shaper.ApplyTitle()
	}
	return a
}

////////////////////////////////////////////////////////////////////////////
// Extending shaper.Shaper

// Shaper extends shaper.Shaper
type Shaper struct {
	*shaper.Shaper
}

// NewShaper makes a new Shaper filter
func NewShaper() *Shaper {
	return &Shaper{Shaper: shaper.NewShaper()}
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

func (sp *Shaper) ApplyToneShaping(a Pinyin) *Shaper {
	sp.AddShaper(func(p string) string {
		if a.truncate == Initials || a.tone == Tone3 {
			// already shortened or no need to change
			return p
		}

		// 替换拼音中的带声调字符
		py := rePhoneticSymbol.ReplaceAllStringFunc(p, func(m string) string {
			symbol, _ := phoneticSymbol[m]
			switch a.tone {
			// 不包含声调
			case Normal:
				// 去掉声调: a1 -> a 匹配使用数字标识声调的字符的正则表达式
				m = regexp.MustCompile("([aeoiuvnm])([1-4])$").
					ReplaceAllString(symbol, "$1")
			case Tone2, Tone1:
				// 返回使用数字标识声调的字符
				m = symbol
			default:
				// 声调在头上
			}
			return m
		})

		if a.tone == Tone1 {
			// 将声调移动到最后. 匹配 Tone2 中标识韵母声调的正则表达式
			py = regexp.MustCompile("^([a-z]+)([1-4])([a-z]*)$").
				ReplaceAllString(py, "$1$3$2")
		}
		return py
	})
	return sp
}

func (sp *Shaper) ApplyTruncate(a Pinyin) *Shaper {
	sp.AddShaper(func(p string) string {
		if a.truncate == Both {
			// 双显风格，返回 全部拼音
			return p
		}

		if a.truncate == FirstLetter {
			// 首字母
			return p[:1]
		}

		// 获取拼音中的声母
		s, y := "", ""
		for _, v := range initialArray {
			if strings.HasPrefix(p, v) {
				s = v
				y = p[len(s):]
				break
			}
		}

		if a.truncate == Initials {
			// 声母风格
			return s
		}

		// 韵母风格

		// 转换为 []rune unicode 编码用于获取第一个拼音字符
		// 因为 string 是 utf-8 编码不方便获取第一个拼音字符
		rs := []rune(p)
		switch string(rs[0]) {
		// 因为鼻音没有声母所以不需要去掉声母部分
		case "ḿ", "ń", "ň", "ǹ":
			return p
		}

		// ǖ 特例 j/q/x/y
		matches := regexp.MustCompile("^(j|q|x|y)(ū|ú|ǔ|ù)$").FindStringSubmatch(p)
		// jū -> jǖ
		if len(matches) == 3 && matches[1] != "" && matches[2] != "" {
			y, _ = finalExceptionsMap[matches[2]]
		}
		matches = regexp.MustCompile("^(j|q|x|y)(u)").FindStringSubmatch(p)
		// yuán -> yván
		if len(matches) == 3 && matches[1] != "" && matches[2] != "" {
			y = "v" + p[2:] // yu -> v
		}

		// 简明整齐的处理声母韵母
		if a.truncate == Finals {
			return y
		}

		// 获取拼音中的韵母
		if s == "" {
			y = handleYW(p)
		}

		return y
	})
	return sp
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
		if !a.polyphone && firstComma > 0 {
			value = value[:firstComma]
		}
		// 多音字模式 (Polyphone), output likes "hang/xing"
		if a.polyphone && firstComma > 0 {
			value = strings.Replace(value, ",", "/", -1)
		}
		py := a.shaper.Process(value)
		if a.truncate == Both {
			// 双显风格
			fmt.Fprintf(pys, "%s(%s)%s", string(r), py, a.Separator)
		} else {
			pys.WriteString(py + a.Separator)
		}
	}
	return pys.String()
}
