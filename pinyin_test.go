package pinyin

import (
	"testing"
)

type pinyinFunc func(string) [][]string

type testCase struct {
	style  Style
	result string
}

func testPinyin(t *testing.T, s string, d []testCase) {
	py := NewPinyin()
	for _, tc := range d {
		py.SetStyle(tc.style)
		v := py.Convert(s)
		if v != tc.result {
			t.Errorf(`Expected "%s", got "%s"`, tc.result, v)
		}
	}
}

func TestPinyin(t *testing.T) {
	hans := "中国人"
	testData := []testCase{
		// default
		{
			Style{Normal, Normal},
			"zhong guo ren ",
		},
		// Tone3
		{
			Style{Tone3, Normal},
			"zhōng guó rén ",
		},
		// Tone2
		{
			Style{Tone2, Normal},
			"zho1ng guo2 re2n ",
		},
		// Tone1
		{
			Style{Tone1, Normal},
			"zhong1 guo2 ren2 ",
		},
		// Initials
		{
			Style{Normal, Initials},
			"zh g r ",
		},
		// FirstLetter
		{
			Style{Normal, FirstLetter},
			"z g r ",
		},
		// Finals
		{
			Style{Normal, Finals},
			"ong uo en ",
		},
		// FinalsTone
		{
			Style{Tone3, Finals},
			"ōng uó én ",
		},
		// FinalsTone2
		{
			Style{Tone2, Finals},
			"o1ng uo2 e2n ",
		},
		// FinalsTone1
		{
			Style{Tone1, Finals},
			"ong1 uo2 en2 ",
		},
	}

	testPinyin(t, hans, testData)

}

func TestFinal(t *testing.T) {
	value := "an"
	v := final("an")
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

type testItem struct {
	hans   string
	style  Style
	result string
}

func testPinyinUpdate(t *testing.T, d []testItem) {
	py := NewPinyin()
	for _, tc := range d {
		py.SetStyle(tc.style)
		v := py.Convert(tc.hans)
		if v != tc.result {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
}

func nnTestUpdated(t *testing.T) {
	testData := []testItem{
		// 误把 yu 放到声母列表了
		{"鱼", Style{Tone2, Normal}, "yu2"},
		{"鱼", Style{Tone1, Normal}, "yu2"},
		{"鱼", Style{Normal, Finals}, "v"},
		{"雨", Style{Tone2, Normal}, "yu3"},
		{"雨", Style{Tone1, Normal}, "yu3"},
		{"雨", Style{Normal, Finals}, "v"},
		{"元", Style{Tone2, Normal}, "yua2n"},
		{"元", Style{Tone1, Normal}, "yuan2"},
		{"元", Style{Normal, Finals}, "van"},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		{"呀", Style{Normal, Initials}, ""},
		{"呀", Style{Tone2, Normal}, "ya"},
		{"呀", Style{Tone1, Normal}, "ya"},
		{"呀", Style{Normal, Finals}, "ia"},
		{"无", Style{Normal, Initials}, ""},
		{"无", Style{Tone2, Normal}, "wu2"},
		{"无", Style{Tone1, Normal}, "wu2"},
		{"无", Style{Normal, Finals}, "u"},
		{"衣", Style{Tone2, Normal}, "yi1"},
		{"衣", Style{Tone1, Normal}, "yi1"},
		{"衣", Style{Normal, Finals}, "i"},
		{"万", Style{Tone2, Normal}, "wa4n"},
		{"万", Style{Tone1, Normal}, "wan4"},
		{"万", Style{Normal, Finals}, "uan"},
		// ju, qu, xu 的韵母应该是 v
		{"具", Style{Tone3, Finals}, "ǜ"},
		{"具", Style{Tone2, Finals}, "v4"},
		{"具", Style{Tone1, Finals}, "v4"},
		{"具", Style{Normal, Finals}, "v"},
		{"取", Style{Tone3, Finals}, "ǚ"},
		{"取", Style{Tone2, Finals}, "v3"},
		{"取", Style{Tone1, Finals}, "v3"},
		{"取", Style{Normal, Finals}, "v"},
		{"徐", Style{Tone3, Finals}, "ǘ"},
		{"徐", Style{Tone2, Finals}, "v2"},
		{"徐", Style{Tone1, Finals}, "v2"},
		{"徐", Style{Normal, Finals}, "v"},
		// # ń
		{"嗯", Style{Normal, Normal}, "n"},
		{"嗯", Style{Tone3, Normal}, "ń"},
		{"嗯", Style{Tone2, Normal}, "n2"},
		{"嗯", Style{Tone1, Normal}, "n2"},
		{"嗯", Style{Normal, Initials}, ""},
		{"嗯", Style{Normal, FirstLetter}, "n"},
		{"嗯", Style{Normal, Finals}, "n"},
		{"嗯", Style{Tone3, Finals}, "ń"},
		{"嗯", Style{Tone2, Finals}, "n2"},
		{"嗯", Style{Tone1, Finals}, "n2"},
		// # ḿ  \u1e3f  U+1E3F
		{"呣", Style{Normal, Normal}, "m"},
		{"呣", Style{Tone3, Normal}, "ḿ"},
		{"呣", Style{Tone2, Normal}, "m2"},
		{"呣", Style{Tone1, Normal}, "m2"},
		{"呣", Style{Normal, Initials}, ""},
		{"呣", Style{Normal, FirstLetter}, "m"},
		{"呣", Style{Normal, Finals}, "m"},
		{"呣", Style{Tone3, Finals}, "ḿ"},
		{"呣", Style{Tone2, Finals}, "m2"},
		{"呣", Style{Tone1, Finals}, "m2"},
		// 去除 0
		{"啊", Style{Tone2, Normal}, "a"},
		{"啊", Style{Tone1, Normal}, "a"},
		{"侵略", Style{Tone2, Normal}, "qi1n lve4"},
		{"侵略", Style{Tone2, Finals}, "i1n ve4"},
		{"侵略", Style{Tone1, Finals}, "in1 ve4"},
	}
	testPinyinUpdate(t, testData)
}
