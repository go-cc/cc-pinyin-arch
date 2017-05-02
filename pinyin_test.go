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
			NewPinyin(Normal, Normal, Separator, false),
			"zhong guo ren ",
		},
		// Tone3
		{
			NewPinyin(Tone3, Normal, Separator, false),
			"zhōng guó rén ",
		},
		// Tone2
		{
			NewPinyin(Tone2, Normal, Separator, false),
			"zho1ng guo2 re2n ",
		},
		// Tone1
		{
			NewPinyin(Tone1, Normal, Separator, false),
			"zhong1 guo2 ren2 ",
		},
		// Initials
		{
			NewPinyin(Normal, Initials, Separator, false),
			"zh g r ",
		},
		// FirstLetter
		{
			NewPinyin(Normal, FirstLetter, Separator, false),
			"z g r ",
		},
		// Finals
		{
			NewPinyin(Normal, Finals, Separator, false),
			"ong uo en ",
		},
		// FinalsTone
		{
			NewPinyin(Tone3, Finals, Separator, false),
			"ōng uó én ",
		},
		// FinalsTone2
		{
			NewPinyin(Tone2, Finals, Separator, false),
			"o1ng uo2 e2n ",
		},
		// FinalsTone1
		{
			NewPinyin(Tone1, Finals, Separator, false),
			"ong1 uo2 en2 ",
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
		if v != tc.result {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
}

func nnTestUpdated(t *testing.T) {
	Separator := " "
	testData := []testItem{
		// 误把 yu 放到声母列表了
		{"鱼", NewPinyin(Tone2, Normal, Separator, false), "yu2"},
		{"鱼", NewPinyin(Tone1, Normal, Separator, false), "yu2"},
		{"鱼", NewPinyin(Normal, Finals, Separator, false), "v"},
		{"雨", NewPinyin(Tone2, Normal, Separator, false), "yu3"},
		{"雨", NewPinyin(Tone1, Normal, Separator, false), "yu3"},
		{"雨", NewPinyin(Normal, Finals, Separator, false), "v"},
		{"元", NewPinyin(Tone2, Normal, Separator, false), "yua2n"},
		{"元", NewPinyin(Tone1, Normal, Separator, false), "yuan2"},
		{"元", NewPinyin(Normal, Finals, Separator, false), "van"},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		{"呀", NewPinyin(Normal, Initials, Separator, false), ""},
		{"呀", NewPinyin(Tone2, Normal, Separator, false), "ya"},
		{"呀", NewPinyin(Tone1, Normal, Separator, false), "ya"},
		{"呀", NewPinyin(Normal, Finals, Separator, false), "ia"},
		{"无", NewPinyin(Normal, Initials, Separator, false), ""},
		{"无", NewPinyin(Tone2, Normal, Separator, false), "wu2"},
		{"无", NewPinyin(Tone1, Normal, Separator, false), "wu2"},
		{"无", NewPinyin(Normal, Finals, Separator, false), "u"},
		{"衣", NewPinyin(Tone2, Normal, Separator, false), "yi1"},
		{"衣", NewPinyin(Tone1, Normal, Separator, false), "yi1"},
		{"衣", NewPinyin(Normal, Finals, Separator, false), "i"},
		{"万", NewPinyin(Tone2, Normal, Separator, false), "wa4n"},
		{"万", NewPinyin(Tone1, Normal, Separator, false), "wan4"},
		{"万", NewPinyin(Normal, Finals, Separator, false), "uan"},
		// ju, qu, xu 的韵母应该是 v
		{"具", NewPinyin(Tone3, Finals, Separator, false), "ǜ"},
		{"具", NewPinyin(Tone2, Finals, Separator, false), "v4"},
		{"具", NewPinyin(Tone1, Finals, Separator, false), "v4"},
		{"具", NewPinyin(Normal, Finals, Separator, false), "v"},
		{"取", NewPinyin(Tone3, Finals, Separator, false), "ǚ"},
		{"取", NewPinyin(Tone2, Finals, Separator, false), "v3"},
		{"取", NewPinyin(Tone1, Finals, Separator, false), "v3"},
		{"取", NewPinyin(Normal, Finals, Separator, false), "v"},
		{"徐", NewPinyin(Tone3, Finals, Separator, false), "ǘ"},
		{"徐", NewPinyin(Tone2, Finals, Separator, false), "v2"},
		{"徐", NewPinyin(Tone1, Finals, Separator, false), "v2"},
		{"徐", NewPinyin(Normal, Finals, Separator, false), "v"},
		// # ń
		{"嗯", NewPinyin(Normal, Normal, Separator, false), "n"},
		{"嗯", NewPinyin(Tone3, Normal, Separator, false), "ń"},
		{"嗯", NewPinyin(Tone2, Normal, Separator, false), "n2"},
		{"嗯", NewPinyin(Tone1, Normal, Separator, false), "n2"},
		{"嗯", NewPinyin(Normal, Initials, Separator, false), ""},
		{"嗯", NewPinyin(Normal, FirstLetter, Separator, false), "n"},
		{"嗯", NewPinyin(Normal, Finals, Separator, false), "n"},
		{"嗯", NewPinyin(Tone3, Finals, Separator, false), "ń"},
		{"嗯", NewPinyin(Tone2, Finals, Separator, false), "n2"},
		{"嗯", NewPinyin(Tone1, Finals, Separator, false), "n2"},
		// # ḿ  \u1e3f  U+1E3F
		{"呣", NewPinyin(Normal, Normal, Separator, false), "m"},
		{"呣", NewPinyin(Tone3, Normal, Separator, false), "ḿ"},
		{"呣", NewPinyin(Tone2, Normal, Separator, false), "m2"},
		{"呣", NewPinyin(Tone1, Normal, Separator, false), "m2"},
		{"呣", NewPinyin(Normal, Initials, Separator, false), ""},
		{"呣", NewPinyin(Normal, FirstLetter, Separator, false), "m"},
		{"呣", NewPinyin(Normal, Finals, Separator, false), "m"},
		{"呣", NewPinyin(Tone3, Finals, Separator, false), "ḿ"},
		{"呣", NewPinyin(Tone2, Finals, Separator, false), "m2"},
		{"呣", NewPinyin(Tone1, Finals, Separator, false), "m2"},
		// 去除 0
		{"啊", NewPinyin(Tone2, Normal, Separator, false), "a"},
		{"啊", NewPinyin(Tone1, Normal, Separator, false), "a"},
		{"侵略", NewPinyin(Tone2, Normal, Separator, false), "qi1n lve4"},
		{"侵略", NewPinyin(Tone2, Finals, Separator, false), "i1n ve4"},
		{"侵略", NewPinyin(Tone1, Finals, Separator, false), "in1 ve4"},
	}
	testPinyinUpdate(t, testData)
}
