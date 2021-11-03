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

type Romanization struct {
	Symbol, Romanji string
}

type Range struct {
	Start, End uint
}


var hir = flag.Bool("hi", false, "Complete Hiragana kana")
var kat = flag.Bool("ka", false, "Complete Katakana kana")

func main() {
	flag.Usage = usage
	flag.Parse()

	ranges := genranges()
	k := genkana(ranges)
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

func genranges() (r []Range) {
	var args []string
	m := map[string]bool{}

	// remove repeated arguments
	for _,v := range flag.Args() {
		if !m[v] {
			args = append(args, v)
			m[v] = true
		}
	}

	if *hir {
		r = append(r, 
			H, HK, HS, HT, HN, HH, HM, HY, HR, HW, Hn, HG, HZ, HD, HB, HP, 
			HKY, HSH, HCH, HNY, HHY, HMY, HRY, HGY, HJ0, HJ1, HBY, HPY,
		)
	} else {
		for _,arg := range args {
			switch arg {
			case "H":
				r = append(r, H)
			case "HK":
				r = append(r, HK)
			case "HS":
				r = append(r, HS)
			case "HT":
				r = append(r, HT)
			case "HN":
				r = append(r, HN)
			case "HH":
				r = append(r, HH)
			case "HM":
				r = append(r, HM)
			case "HY":
				r = append(r, HY)
			case "HR":
				r = append(r, HR)
			case "HW":
				r = append(r, HW)
			case "Hn":
				r = append(r, Hn)
			case "HG":
				r = append(r, HG)
			case "HZ":
				r = append(r, HZ)
			case "HD":
				r = append(r, HD)
			case "HB":
				r = append(r, HB)
			case "HP":
				r = append(r, HP)
			case "HKY":
				r = append(r, HKY)
			case "HSH":
				r = append(r, HSH)
			case "HCH":
				r = append(r, HCH)
			case "HNY":
				r = append(r, HNY)
			case "HHY":
				r = append(r, HHY)
			case "HMY":
				r = append(r, HMY)
			case "HRY":
				r = append(r, HRY)
			case "HGY":
				r = append(r, HGY)
			case "HJ0":
				r = append(r, HJ0)
			case "HJ1":
				r = append(r, HJ1)
			case "HBY":
				r = append(r, HBY)
			case "HPY":
				r = append(r, HPY)
			}
		}
	}

	if *kat {
		r = append(r, 
			K, KK, KS, KT, KN, KH, KM, KY, KR, KW, Kn, KG, KZ, KD, KB, KP, 
			KKY, KSH, KCH, KNY, KHY, KMY, KRY, KGY, KJ0, KJ1, KBY, KPY,
		)
	} else {
		for _,arg := range args {
			switch arg {
			case "K":
				r = append(r, K)
			case "KK":
				r = append(r, KK)
			case "KS":
				r = append(r, KS)
			case "KT":
				r = append(r, KT)
			case "KN":
				r = append(r, KN)
			case "KH":
				r = append(r, KH)
			case "KM":
				r = append(r, KM)
			case "KY":
				r = append(r, KY)
			case "KR":
				r = append(r, KR)
			case "KW":
				r = append(r, KW)
			case "Kn":
				r = append(r, Kn)
			case "KG":
				r = append(r, KG)
			case "KZ":
				r = append(r, KZ)
			case "KD":
				r = append(r, KD)
			case "KB":
				r = append(r, KB)
			case "KP":
				r = append(r, KP)
			case "KKY":
				r = append(r, KKY)
			case "KSH":
				r = append(r, KSH)
			case "KCH":
				r = append(r, KCH)
			case "KNY":
				r = append(r, KNY)
			case "KHY":
				r = append(r, KHY)
			case "KMY":
				r = append(r, KMY)
			case "KRY":
				r = append(r, KRY)
			case "KGY":
				r = append(r, KGY)
			case "KJ0":
				r = append(r, KJ0)
			case "KJ1":
				r = append(r, KJ1)
			case "KBY":
				r = append(r, KBY)
			case "KPY":
				r = append(r, KPY)
			}
		}
	}

	return
}

func genkana(ranges []Range) []Romanization {
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
	
	if ranges == nil || len(ranges) == 0 {
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

func usage() {
	out := flag.CommandLine.Output()

	fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprint(out, "  Hiragana:\n")
	tblprt(out, "H ", genkana([]Range{H}))
	tblprt(out, "HK", genkana([]Range{HK}))
	tblprt(out, "HS", genkana([]Range{HS}))
	tblprt(out, "HT", genkana([]Range{HT}))
	tblprt(out, "HN", genkana([]Range{HN}))
	tblprt(out, "HH", genkana([]Range{HH}))
	tblprt(out, "HM", genkana([]Range{HM}))
	tblprt(out, "HY", genkana([]Range{HY}))
	tblprt(out, "HR", genkana([]Range{HR}))
	tblprt(out, "HW", genkana([]Range{HW}))
	tblprt(out, "Hn", genkana([]Range{Hn}))
	tblprt(out, "HG", genkana([]Range{HG}))
	tblprt(out, "HZ", genkana([]Range{HZ}))
	tblprt(out, "HD", genkana([]Range{HD}))
	tblprt(out, "HB", genkana([]Range{HB}))
	tblprt(out, "HP", genkana([]Range{HP}))

	fmt.Fprint(out, "  Hiragana Combinations:\n")
	tblprt(out, "HKY", genkana([]Range{HKY}))
	tblprt(out, "HSH", genkana([]Range{HSH}))
	tblprt(out, "HCH", genkana([]Range{HCH}))
	tblprt(out, "HNY", genkana([]Range{HNY}))
	tblprt(out, "HMY", genkana([]Range{HMY}))
	tblprt(out, "HRY", genkana([]Range{HRY}))
	tblprt(out, "HGY", genkana([]Range{HGY}))
	tblprt(out, "HJ0", genkana([]Range{HJ0}))
	tblprt(out, "HJ1", genkana([]Range{HJ1}))
	tblprt(out, "HBY", genkana([]Range{HBY}))
	tblprt(out, "HPY", genkana([]Range{HPY}))

	fmt.Fprint(out, "  Katakana:\n")
	tblprt(out, "K ", genkana([]Range{K}))
	tblprt(out, "KK", genkana([]Range{KK}))
	tblprt(out, "KS", genkana([]Range{KS}))
	tblprt(out, "KT", genkana([]Range{KT}))
	tblprt(out, "KN", genkana([]Range{KN}))
	tblprt(out, "KH", genkana([]Range{KH}))
	tblprt(out, "KM", genkana([]Range{KM}))
	tblprt(out, "KY", genkana([]Range{KY}))
	tblprt(out, "KR", genkana([]Range{KR}))
	tblprt(out, "KW", genkana([]Range{KW}))
	tblprt(out, "Kn", genkana([]Range{Kn}))
	tblprt(out, "KG", genkana([]Range{KG}))
	tblprt(out, "KZ", genkana([]Range{KZ}))
	tblprt(out, "KD", genkana([]Range{KD}))
	tblprt(out, "KB", genkana([]Range{KB}))
	tblprt(out, "KP", genkana([]Range{KP}))

	fmt.Fprint(out, "  Katakana Combinations:\n")
	tblprt(out, "KKY", genkana([]Range{KKY}))
	tblprt(out, "KSH", genkana([]Range{KSH}))
	tblprt(out, "KCH", genkana([]Range{KCH}))
	tblprt(out, "KNY", genkana([]Range{KNY}))
	tblprt(out, "KMY", genkana([]Range{KMY}))
	tblprt(out, "KRY", genkana([]Range{KRY}))
	tblprt(out, "KGY", genkana([]Range{KGY}))
	tblprt(out, "KJ0", genkana([]Range{KJ0}))
	tblprt(out, "KJ1", genkana([]Range{KJ1}))
	tblprt(out, "KBY", genkana([]Range{KBY}))
	tblprt(out, "KPY", genkana([]Range{KPY}))
}

func tblprt(out io.Writer, header string, line []Romanization) {
	fmt.Fprintf(out, "\t%s:", header)
	for _,v := range line {
		fmt.Fprintf(out, " %s", v.Symbol)
	}
	fmt.Fprintln(out)
}

/* Hiragana */
var H Range = Range {
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
var K Range = Range {
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
