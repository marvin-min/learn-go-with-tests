package intro_to_property_based_tests

import (
	"strings"
)

/**
  * Ⅰ－1、Ⅱ－2、Ⅲ－3、Ⅳ－4、Ⅴ－5、Ⅵ－6、Ⅶ－7、Ⅷ－8、Ⅸ－9
	* Ⅹ－10、Ⅺ－11、Ⅻ－12、XIII－13、XIV－14、XV－15、XVI－16、XVII－17、XVIII－18、XIX－19、XX－20、XXI－21、XXII－22、XXIX－29、XXX－30、XXXIV－34、XXXV－35、XXXIX－39、XL－40、XLV－45、XLIX（IL）－49、L－50、LI－51、LV－55、LX－60、LXV－65、LXXX－80、XC－90、XCIII－93、XCV－95、XCVIII－98、XCIX（IC）－99
	* C－100、CC－200、CCC－300、CD－400、CDXC（XD）－490、CDXCV－495、CDXCIX－499、D－500、DC－600、DCC－700、DCCC－800、CM－900、CMXCIX－999
	* M－1000、MC－1100、MCD－1400、MD－1500、MDC－1600、MDCLXVI－1666、MDCCCLXXXVIII－1888、MDCCCXCIX（MDCCCIC）－1899、MCM－1900、MCMLXXVI－1976、MCMLXXXIV－1984、MCMXC（MXM）－1990、MM－2000、MMMCMXCIX（MMMIM）－3999
	**/

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 9:
	// 		result.WriteString("X")
	// 		arabic -= 10
	// 	case arabic > 8:
	// 		result.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		result.WriteString("V")
	// 		arabic -= 5
	// 	case arabic > 3:
	// 		result.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabic--
	// 	}
	// }
	// for i := arabic; i > 0; i-- {
	// 	if i == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}
	// 	result.WriteString("I")
	// }
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0
	for _, numeral := range allRomanNumerals {
		// fmt.Println(roman, " VS ", numeral.Symbol, "---->", strings.HasPrefix(roman, numeral.Symbol))
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
			// fmt.Println(roman)
		}
	}
	return arabic
}
