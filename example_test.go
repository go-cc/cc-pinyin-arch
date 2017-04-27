package pinyin_test

import (
	"fmt"

	pinyin "github.com/go-cc/cc-pinyin"
)

// for standalone test, change package to main and the next func def to,
// func main() {
func Example_output() {
	s := `名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。`
	_ = s
	hans := "中国人的中国银行很行"

	// 默认
	a := pinyin.NewPinyin()
	a.Separator = "-"
	fmt.Println(a.Convert(hans))

	// 包含声调
	a.SetStyle(pinyin.Tone)
	fmt.Println(a.Convert(hans))

	// 声调用数字表示
	a.SetStyle(pinyin.Tone2)
	fmt.Println(a.Convert(hans))

	// 声调在拼音后用数字表示
	a.SetStyle(pinyin.Tone3)
	fmt.Println(a.Convert(hans))

	// 开启多音字模式
	a.Heteronym = true
	fmt.Println(a.Convert(hans))
	a.SetStyle(pinyin.Tone)
	fmt.Println(a.Convert(hans))

	// Output:
	// MingZhu：《HongLouMeng》〖Qing〗CaoXueQin Zhu、GaoZuo Xu／『RenMinWenXue』ChuBanShe／1996—9Yue30Ri／59.70【Yuan】，《SanGuoYanYi》〖Ming〗LuoGuanZhong。
	// ming-zhu-：《hong-lou-meng-》〖qing-〗cao-xue-qin- zhu-、gao-zuo- xu-／『ren-min-wen-xue-』chu-ban-she-／1996—9yue-30ri-／59.70【yuan-】，《san-guo-yan-yi-》〖ming-〗luo-guan-zhong-。

	// Output:
	// zhong-guo-ren-de-zhong-guo-yin-xing-hen-xing-
	// zhōng-guó-rén-de-zhōng-guó-yín-xíng-hěn-xíng-
	// zho1ng-guo2-re2n-de-zho1ng-guo2-yi2n-xi2ng-he3n-xi2ng-
	// zhong1-guo2-ren2-de-zhong1-guo2-yin2-xing2-hen3-xing2-
	// zho1ng/zho4ng-guo2-ren2-de/di4/di2-zho1ng/zho4ng-guo2-yin2-xi2ng/ha2ng/xi4ng/ha4ng/he2ng-hen3-xi2ng/ha2ng/xi4ng/ha4ng/he2ng-
	// zhōng/zhòng-guó-rén-de/dì/dí-zhōng/zhòng-guó-yín-xíng/háng/xìng/hàng/héng-hěn-xíng/háng/xìng/hàng/héng-

}
