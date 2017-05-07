package pinyin

import (
	"testing"
)

type pinyinFunc func(string) [][]string

type testCase struct {
	a      Pinyin
	result string
}

func testPinyin(t *testing.T, s string, d []testCase) {
	for _, tc := range d {
		v := tc.a.Convert(s)
		if v != tc.result {
			t.Errorf(`Expected "%s", got "%s"`, tc.result, v)
		}
	}
}

func TestPinyin(t *testing.T) {
	hans := "中国人"
	Separator := " "
	testData := []testCase{
		// default
		{
			NewPinyin(Normal, Normal, Separator, false, false),
			"zhong guo ren ",
		},
		{
			NewPinyin(Normal, Normal, Separator, false, true),
			"Zhong Guo Ren ",
		},
		// Tone3
		{
			NewPinyin(Tone3, Normal, Separator, false, false),
			"zhōng guó rén ",
		},
		{
			NewPinyin(Tone3, Normal, Separator, false, true),
			"Zhōng Guó Rén ",
		},
		// Tone2
		{
			NewPinyin(Tone2, Normal, Separator, false, false),
			"zho1ng guo2 re2n ",
		},
		{
			NewPinyin(Tone2, Normal, Separator, false, true),
			"Zho1ng Guo2 Re2n ",
		},
		// Tone1
		{
			NewPinyin(Tone1, Normal, Separator, false, false),
			"zhong1 guo2 ren2 ",
		},
		{
			NewPinyin(Tone1, Normal, Separator, false, true),
			"Zhong1 Guo2 Ren2 ",
		},
		// Initials
		{
			NewPinyin(Normal, Initials, Separator, false, false),
			"zh g r ",
		},
		{
			NewPinyin(Normal, Initials, Separator, false, true),
			"Zh G R ",
		},
		// FirstLetter
		{
			NewPinyin(Normal, FirstLetter, Separator, false, false),
			"z g r ",
		},
		{
			NewPinyin(Normal, FirstLetter, Separator, false, true),
			"Z G R ",
		},
		// Finals
		{
			NewPinyin(Normal, Finals, Separator, false, false),
			"ong uo en ",
		},
		{
			NewPinyin(Normal, Finals, Separator, false, true),
			"Ong Uo En ",
		},
		// FinalsTone
		{
			NewPinyin(Tone3, Finals, Separator, false, false),
			"ōng uó én ",
		},
		{
			NewPinyin(Tone3, Finals, Separator, false, true),
			"Ōng Uó Én ",
		},
		// FinalsTone2
		{
			NewPinyin(Tone2, Finals, Separator, false, false),
			"o1ng uo2 e2n ",
		},
		{
			NewPinyin(Tone2, Finals, Separator, false, true),
			"O1ng Uo2 E2n ",
		},
		// FinalsTone1
		{
			NewPinyin(Tone1, Finals, Separator, false, false),
			"ong1 uo2 en2 ",
		},
		{
			NewPinyin(Tone1, Finals, Separator, false, true),
			"Ong1 Uo2 En2 ",
		},
		{
			NewPinyin(Tone3, Both, Separator, false, true),
			"中(Zhōng) 国(Guó) 人(Rén) ",
		},
	}

	testPinyin(t, hans, testData)

}

type testItem struct {
	hans   string
	a      Pinyin
	result string
}

func testPinyinUpdate(t *testing.T, d []testItem) {
	for _, tc := range d {
		v := tc.a.Convert(tc.hans)
		//t.Log(v)
		if v != tc.result {
			t.Errorf("'%s' expects '%s', got '%s'", tc.hans, tc.result, v)
		}
	}
}

func TestUpdated(t *testing.T) {
	Separator := " "
	testData := []testItem{
		// 误把 yu 放到声母列表了
		{"鱼", NewPinyin(Tone2, Normal, Separator, false, false), "yu2 "},
		{"鱼", NewPinyin(Tone1, Normal, Separator, false, false), "yu2 "},
		{"鱼", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "v "},
		{"雨", NewPinyin(Tone2, Normal, Separator, false, false), "yu3 "},
		{"雨", NewPinyin(Tone1, Normal, Separator, false, false), "yu3 "},
		{"雨", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "v "},
		{"元", NewPinyin(Tone3, Normal, Separator, false, false), "yuán "},
		{"元", NewPinyin(Tone2, Normal, Separator, false, false), "yua2n "},
		{"元", NewPinyin(Tone1, Normal, Separator, false, false), "yuan2 "},
		{"元", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "van "},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		// {"呀", NewPinyin(Normal, Initials, Separator, false, false), " "},
		{"呀", NewPinyin(Tone2, Normal, Separator, false, false), "ya "},
		{"呀", NewPinyin(Tone1, Normal, Separator, false, false), "ya "},
		// {"呀", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "ia "},
		// {"无", NewPinyin(Normal, Initials, Separator, false, false), " "},
		{"无", NewPinyin(Tone2, Normal, Separator, false, false), "wu2 "},
		{"无", NewPinyin(Tone1, Normal, Separator, false, false), "wu2 "},
		{"无", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "u "},
		{"衣", NewPinyin(Tone2, Normal, Separator, false, false), "yi1 "},
		{"衣", NewPinyin(Tone1, Normal, Separator, false, false), "yi1 "},
		{"衣", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "i "},
		{"万", NewPinyin(Tone2, Normal, Separator, false, false), "wa4n "},
		{"万", NewPinyin(Tone1, Normal, Separator, false, false), "wan4 "},
		// {"万", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "uan "},
		// ju, qu, xu 的韵母应该是 v
		{"具", NewPinyin(Tone3, ZeroConsonant, Separator, false, false), "ǜ "},
		{"具", NewPinyin(Tone2, ZeroConsonant, Separator, false, false), "v4 "},
		{"具", NewPinyin(Tone1, ZeroConsonant, Separator, false, false), "v4 "},
		{"具", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "v "},
		{"取", NewPinyin(Tone3, ZeroConsonant, Separator, false, false), "ǚ "},
		{"取", NewPinyin(Tone2, ZeroConsonant, Separator, false, false), "v3 "},
		{"取", NewPinyin(Tone1, ZeroConsonant, Separator, false, false), "v3 "},
		{"取", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "v "},
		{"徐", NewPinyin(Tone3, ZeroConsonant, Separator, false, false), "ǘ "},
		{"徐", NewPinyin(Tone2, ZeroConsonant, Separator, false, false), "v2 "},
		{"徐", NewPinyin(Tone1, ZeroConsonant, Separator, false, false), "v2 "},
		{"徐", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "v "},
		// # ń
		{"嗯", NewPinyin(Normal, Normal, Separator, false, false), "n "},
		{"嗯", NewPinyin(Tone3, Normal, Separator, false, false), "ń "},
		{"嗯", NewPinyin(Tone2, Normal, Separator, false, false), "n2 "},
		{"嗯", NewPinyin(Tone1, Normal, Separator, false, false), "n2 "},
		{"嗯", NewPinyin(Normal, Initials, Separator, false, false), " "},
		// {"嗯", NewPinyin(Normal, FirstLetter, Separator, false, false), "n "},
		{"嗯", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "n "},
		{"嗯", NewPinyin(Tone3, ZeroConsonant, Separator, false, false), "ń "},
		{"嗯", NewPinyin(Tone2, ZeroConsonant, Separator, false, false), "n2 "},
		{"嗯", NewPinyin(Tone1, ZeroConsonant, Separator, false, false), "n2 "},
		// # ḿ  \u1e3f  U+1E3F
		// {"呣", NewPinyin(Normal, Normal, Separator, false, false), "m "},
		{"呣", NewPinyin(Tone3, Normal, Separator, false, false), "ḿ "},
		{"呣", NewPinyin(Tone2, Normal, Separator, false, false), "m2 "},
		{"呣", NewPinyin(Tone1, Normal, Separator, false, false), "m2 "},
		{"呣", NewPinyin(Normal, Initials, Separator, false, false), " "},
		// {"呣", NewPinyin(Normal, FirstLetter, Separator, false, false), "m "},
		{"呣", NewPinyin(Normal, ZeroConsonant, Separator, false, false), "m "},
		{"呣", NewPinyin(Tone3, ZeroConsonant, Separator, false, false), "ḿ "},
		{"呣", NewPinyin(Tone2, ZeroConsonant, Separator, false, false), "m2 "},
		{"呣", NewPinyin(Tone1, ZeroConsonant, Separator, false, false), "m2 "},
		// 去除 0
		{"啊", NewPinyin(Tone2, Normal, Separator, false, false), "a "},
		{"啊", NewPinyin(Tone1, Normal, Separator, false, false), "a "},
		{"侵略", NewPinyin(Tone2, Normal, Separator, false, false), "qi1n lve4 "},
		{"侵略", NewPinyin(Tone2, ZeroConsonant, Separator, false, false), "i1n ve4 "},
		{"侵略", NewPinyin(Tone1, ZeroConsonant, Separator, false, false), "in1 ve4 "},
		{"语文", NewPinyin(Tone1, Initials, Separator, false, false), "y w "},
		{"语文", NewPinyin(Tone3, Initials, Separator, false, false), "y w "},
		{"语文", NewPinyin(Tone1, Finals, Separator, false, false), "v3 en2 "},
		{"语文", NewPinyin(Tone3, Finals, Separator, false, false), "ǚ én "},
	}
	testPinyinUpdate(t, testData)
}
