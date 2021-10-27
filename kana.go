package main

import (
	"fmt"
	"os"
	"io"
	"time"
	"flag"
	"sort"
	"bufio"
	"strings"
	"math/rand"
)

var hir = flag.Bool("h", false, "Complete Hiragana kana")
var kat = flag.Bool("k", false, "Complete Katakana kana")

func main() {
	flag.Parse()

	var ranges []Range	
	if *hir {
		ranges = append(ranges, Hempty, HK, HS, HT, HN, HH, HM, HY, HR, HW, Hn, HG, HZ, HD, HB, HP, HKY, HSH, HCH, HNY, HHY, HMY, HRY, HGY, HJ0, HJ1, HBY, HPY)
	}
	if *kat {
		ranges = append(ranges, Kempty, KK, KS, KT, KN, KH, KM, KY, KR, KW, Kn, KG, KZ, KD, KB, KP, KKY, KSH, KCH, KNY, KHY, KMY, KRY, KGY, KJ0, KJ1, KBY, KPY)
	}

	k := gen(ranges)
	r := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	OUT:
	for {
		var insymbol string
		chosen := choose(k)
		
		for insymbol != chosen.Romanji {
			fmt.Printf("%s:\t", chosen.Symbol)

			line, err := r.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break OUT
				}
				perr("reading line", err)
			}

			insymbol = strings.TrimRight(line, "\n")
			if insymbol != chosen.Romanji {
				fmt.Printf("(%s)\n", chosen.Romanji)
			}
		}
	}
}

func choose(k []Romanization) Romanization {
	return k[rand.Intn(len(k))]
}

func perr(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err)
	os.Exit(1)
}

func gen(ranges []Range) []Romanization {
	kana := []Romanization {
		{ "あ", "a", },
		{ "い", "i", },
		{ "う", "u", },
		{ "え", "e", },
		{ "お", "o", },
		{ "か", "ka", },
		{ "き", "ki", },
		{ "く", "ku", },
		{ "け", "ke", },
		{ "こ", "ko", },
		{ "さ", "sa", },
		{ "し", "shi", },
		{ "す", "su", },
		{ "せ", "se", },
		{ "そ", "so", },
		{ "た", "ta", },
		{ "ち", "chi", },
		{ "つ", "tsu", },
		{ "て", "te", },
		{ "と", "to", },
		{ "な", "na", },
		{ "に", "ni", },
		{ "ぬ", "nu", },
		{ "ね", "ne", },
		{ "の", "no", },
		{ "は", "ha", },
		{ "ひ", "hi", },
		{ "ふ", "fu", },
		{ "へ", "he", },
		{ "ほ", "ho", },
		{ "ま", "ma", },
		{ "み", "mi", },
		{ "む", "mu", },
		{ "め", "me", },
		{ "も", "mo", },
		{ "や", "ya", },
		{ "ゆ", "yu", },
		{ "よ", "yo", },
		{ "ら", "ra", },
		{ "り", "ri", },
		{ "る", "ru", },
		{ "れ", "re", },
		{ "ろ", "ro", },
		{ "わ", "wa", },
		{ "を", "o", },
		{ "ん", "n", },
		{ "が", "ga", },
		{ "ぎ", "gi", },
		{ "ぐ", "gu", },
		{ "げ", "ge", },
		{ "ご", "go", },
		{ "ざ", "za", },
		{ "じ", "ji", },
		{ "ず", "zu", },
		{ "ぜ", "ze", },
		{ "ぞ", "zo", },
		{ "だ", "da", },
		{ "ぢ", "ji", },
		{ "づ", "zu", },
		{ "で", "de", },
		{ "ど", "do", },
		{ "ば", "ba", },
		{ "び", "bi", },
		{ "ぶ", "bu", },
		{ "べ", "be", },
		{ "ぼ", "bo", },
		{ "ぱ", "pa", },
		{ "ぴ", "pi", },
		{ "ぷ", "pu", },
		{ "ぺ", "pe", },
		{ "ぽ", "po", },
		{ "きゃ", "kya", },
		{ "きゅ", "kyu", },
		{ "きょ", "kyo", },
		{ "しゃ", "sha", },
		{ "しゅ", "shu", },
		{ "しょ", "sho", },
		{ "ちゃ", "cha", },
		{ "ちゅ", "chu", },
		{ "ちょ", "cho", },
		{ "にゃ", "nya", },
		{ "にゅ", "nyu", },
		{ "にょ", "nyo", },
		{ "ひゃ", "hya", },
		{ "ひゅ", "hyu", },
		{ "ひょ", "hyo", },
		{ "みゃ", "mya", },
		{ "みゅ", "myu", },
		{ "みょ", "myo", },
		{ "りゃ", "rya", },
		{ "りゅ", "ryu", },
		{ "りょ", "ryo", },
		{ "ぎゃ", "gya", },
		{ "ぎゅ", "gyu", },
		{ "ぎょ", "gyo", },
		{ "じゃ", "ja", },
		{ "じゅ", "ju", },
		{ "じょ", "jo", },
		{ "ぢゃ", "ja", },
		{ "ぢゅ", "ju", },
		{ "ぢょ", "jo", },
		{ "びゃ", "bya", },
		{ "びゅ", "byu", },
		{ "びょ", "byo", },
		{ "ぴゃ", "pya", },
		{ "ぴゅ", "pyu", },
		{ "ぴょ", "pyo", },
		{ "ア", "a", },
		{ "イ", "i", },
		{ "ウ", "u", },
		{ "エ", "e", },
		{ "オ", "o", },
		{ "カ", "ka", },
		{ "キ", "ki", },
		{ "ク", "ku", },
		{ "ケ", "ke", },
		{ "コ", "ko", },
		{ "サ", "sa", },
		{ "シ", "shi", },
		{ "ス", "su", },
		{ "セ", "se", },
		{ "ソ", "so", },
		{ "タ", "ta", },
		{ "チ", "chi", },
		{ "ツ", "tsu", },
		{ "テ", "te", },
		{ "ト", "to", },
		{ "ナ", "na", },
		{ "ニ", "ni", },
		{ "ヌ", "nu", },
		{ "ネ", "ne", },
		{ "ノ", "no", },
		{ "ハ", "ha", },
		{ "ヒ", "hi", },
		{ "フ", "fu", },
		{ "ヘ", "he", },
		{ "ホ", "ho", },
		{ "マ", "ma", },
		{ "ミ", "mi", },
		{ "ム", "mu", },
		{ "メ", "me", },
		{ "モ", "mo", },
		{ "ヤ", "ya", },
		{ "ユ", "yu", },
		{ "ヨ", "yo", },
		{ "ラ", "ra", },
		{ "リ", "ri", },
		{ "ル", "ru", },
		{ "レ", "re", },
		{ "ロ", "ro", },
		{ "ワ", "wa", },
		{ "ヲ", "o", },
		{ "ン", "n", },
		{ "ガ", "ga", },
		{ "ギ", "gi", },
		{ "グ", "gu", },
		{ "ゲ", "ge", },
		{ "ゴ", "go", },
		{ "ザ", "za", },
		{ "ジ", "ji", },
		{ "ズ", "zu", },
		{ "ゼ", "ze", },
		{ "ゾ", "zo", },
		{ "ダ", "da", },
		{ "ヂ", "ji", },
		{ "ヅ", "zu", },
		{ "デ", "de", },
		{ "ド", "do", },
		{ "バ", "ba", },
		{ "ビ", "bi", },
		{ "ブ", "bu", },
		{ "ベ", "be", },
		{ "ボ", "bo", },
		{ "パ", "pa", },
		{ "ピ", "pi", },
		{ "プ", "pu", },
		{ "ペ", "pe", },
		{ "ポ", "po", },
		{ "キャ", "kya", },
		{ "キュ", "kyu", },
		{ "キョ", "kyo", },
		{ "シャ", "sha", },
		{ "シャ", "shu", },
		{ "ショ", "sho", },
		{ "チャ", "cha", },
		{ "チュ", "chu", },
		{ "チョ", "cho", },
		{ "ニャ", "nya", },
		{ "ニュ", "nyu", },
		{ "ニョ", "nyo", },
		{ "ヒャ", "hya", },
		{ "ヒュ", "hyu", },
		{ "ヒョ", "hyo", },
		{ "ミャ", "mya", },
		{ "ミュ", "myu", },
		{ "ミョ", "myo", },
		{ "リャ", "rya", },
		{ "リュ", "ryu", },
		{ "リョ", "ryo", },
		{ "ギャ", "gya", },
		{ "ギュ", "gyu", },
		{ "ギョ", "gyo", },
		{ "ジャ", "ja", },
		{ "ジュ", "ju", },
		{ "ジョ", "jo", },
		{ "ヂャ", "ja", },
		{ "ヂュ", "ju", },
		{ "ヂョ", "jo", },
		{ "ビャ", "bya", },
		{ "ビュ", "byu", },
		{ "ビョ", "byo", },
		{ "ピャ", "pya", },
		{ "ピュ", "pyu", },
		{ "ピョ", "pyo", },
	}
	
	if ranges == nil {
		return kana
	}

	var k []Romanization

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	for _,r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			k = append(k, kana[i])
		}
	}

	return k
}

type Romanization struct {
	Symbol, Romanji string
}

type Range struct {
	Start, End uint
}

/* Hiragana */
var Hempty Range = Range {
	0, 4,
}
var HK Range = Range {
	5, 9,
}
var HS Range = Range {
	10, 14,
}
var HT Range = Range {
	15, 19,
}
var HN Range = Range {
	20, 24,
}
var HH Range = Range {
	25, 29,
}
var HM Range = Range {
	30, 34,
}
var HY Range = Range {
	35, 37,
}
var HR Range = Range {
	38, 42,
}
var HW Range = Range {
	43, 44,
}
var Hn Range = Range {
	45, 45,
}
var HG Range = Range {
	46, 50,
}
var HZ Range = Range {
	51, 55,
}
var HD Range = Range {
	56, 60,
}
var HB Range = Range {
	61, 65,
}
var HP Range = Range {
	66, 70,
}
/* Hiragana Combinations */
var HKY Range = Range {
	71, 73,
}
var HSH Range = Range {
	74, 76,
}
var HCH Range = Range {
	77, 79,
}
var HNY Range = Range {
	80, 82,
}
var HHY Range = Range {
	83, 85,
}
var HMY Range = Range {
	86, 88,
}
var HRY Range = Range {
	89, 91,
}
var HGY Range = Range {
	92, 94,
}
var HJ0 Range = Range {
	95, 97,
}
var HJ1 Range = Range {
	98, 100,
}
var HBY Range = Range {
	101, 103,
}
var HPY Range = Range {
	104, 106,
}
/* Katakana */
var Kempty Range = Range {
	107, 111,
}
var KK Range = Range {
	112, 116,
}
var KS Range = Range {
	117, 121,
}
var KT Range = Range {
	122, 126,
}
var KN Range = Range {
	127, 131,
}
var KH Range = Range {
	132, 136,
}
var KM Range = Range {
	137, 141,
}
var KY Range = Range {
	142, 144,
}
var KR Range = Range {
	145, 149,
}
var KW Range = Range {
	150, 151,
}
var Kn Range = Range {
	152, 152,
}
var KG Range = Range {
	153, 157,
}
var KZ Range = Range {
	158, 162,
}
var KD Range = Range {
	163, 167,
}
var KB Range = Range {
	168, 172,
}
var KP Range = Range {
	173, 177,
}
var KKY Range = Range {
	178, 180,
}
var KSH Range = Range {
	181, 183,
}
var KCH Range = Range {
	184, 186,
}
var KNY Range = Range {
	187, 189,
}
var KHY Range = Range {
	190, 192,
}
var KMY Range = Range {
	193, 195,
}
var KRY Range = Range {
	196, 198,
}
var KGY Range = Range {
	199, 201,
}
var KJ0 Range = Range {
	202, 204,
}
var KJ1 Range = Range {
	205, 207,
}
var KBY Range = Range {
	208, 210,
}
var KPY Range = Range {
	211, 213,
}
