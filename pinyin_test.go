package pinyin

import (
	"testing"
)

type pinyinFunc func(string) [][]string

type testCase struct {
	style  int
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
			Normal,
			"zhong guo ren ",
		},
		// Tone
		{
			Tone,
			"zhōng guó rén ",
		},
		// Tone2
		{
			Tone2,
			"zho1ng guo2 re2n ",
		},
		// Tone3
		{
			Tone3,
			"zhong1 guo2 ren2 ",
		},
		// Initials
		{
			Initials,
			"zh g r ",
		},
		// FirstLetter
		{
			FirstLetter,
			"z g r ",
		},
		// Finals
		{
			Finals,
			"ong uo en ",
		},
		// FinalsTone
		{
			FinalsTone,
			"ōng uó én ",
		},
		// FinalsTone2
		{
			FinalsTone2,
			"o1ng uo2 e2n ",
		},
		// FinalsTone3
		{
			FinalsTone3,
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
	style  int
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
		{"鱼", Tone2, "yu2"},
		{"鱼", Tone3, "yu2"},
		{"鱼", Finals, "v"},
		{"雨", Tone2, "yu3"},
		{"雨", Tone3, "yu3"},
		{"雨", Finals, "v"},
		{"元", Tone2, "yua2n"},
		{"元", Tone3, "yuan2"},
		{"元", Finals, "van"},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		{"呀", Initials, ""},
		{"呀", Tone2, "ya"},
		{"呀", Tone3, "ya"},
		{"呀", Finals, "ia"},
		{"无", Initials, ""},
		{"无", Tone2, "wu2"},
		{"无", Tone3, "wu2"},
		{"无", Finals, "u"},
		{"衣", Tone2, "yi1"},
		{"衣", Tone3, "yi1"},
		{"衣", Finals, "i"},
		{"万", Tone2, "wa4n"},
		{"万", Tone3, "wan4"},
		{"万", Finals, "uan"},
		// ju, qu, xu 的韵母应该是 v
		{"具", FinalsTone, "ǜ"},
		{"具", FinalsTone2, "v4"},
		{"具", FinalsTone3, "v4"},
		{"具", Finals, "v"},
		{"取", FinalsTone, "ǚ"},
		{"取", FinalsTone2, "v3"},
		{"取", FinalsTone3, "v3"},
		{"取", Finals, "v"},
		{"徐", FinalsTone, "ǘ"},
		{"徐", FinalsTone2, "v2"},
		{"徐", FinalsTone3, "v2"},
		{"徐", Finals, "v"},
		// # ń
		{"嗯", Normal, "n"},
		{"嗯", Tone, "ń"},
		{"嗯", Tone2, "n2"},
		{"嗯", Tone3, "n2"},
		{"嗯", Initials, ""},
		{"嗯", FirstLetter, "n"},
		{"嗯", Finals, "n"},
		{"嗯", FinalsTone, "ń"},
		{"嗯", FinalsTone2, "n2"},
		{"嗯", FinalsTone3, "n2"},
		// # ḿ  \u1e3f  U+1E3F
		{"呣", Normal, "m"},
		{"呣", Tone, "ḿ"},
		{"呣", Tone2, "m2"},
		{"呣", Tone3, "m2"},
		{"呣", Initials, ""},
		{"呣", FirstLetter, "m"},
		{"呣", Finals, "m"},
		{"呣", FinalsTone, "ḿ"},
		{"呣", FinalsTone2, "m2"},
		{"呣", FinalsTone3, "m2"},
		// 去除 0
		{"啊", Tone2, "a"},
		{"啊", Tone3, "a"},
		{"侵略", Tone2, "qi1n lve4"},
		{"侵略", FinalsTone2, "i1n ve4"},
		{"侵略", FinalsTone3, "in1 ve4"},
	}
	testPinyinUpdate(t, testData)
}
