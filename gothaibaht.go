package gothaibaht

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

// ToText return the thai baht as text.
// This will floor the decimal if the specified number of 2 decimal is exceeded.
func ToText(num float64) string {
	numStr := fmt.Sprintf("%.2f", math.Floor(num*100)/100)
	text := convertNumToText(numStr)
	return text
}

var digits = map[int]string{
	0: "ศูนย์",
	1: "หนึ่ง",
	2: "สอง",
	3: "สาม",
	4: "สี่",
	5: "ห้า",
	6: "หก",
	7: "เจ็ด",
	8: "แปด",
	9: "เก้า",
}

var currency = "บาท"

var suffixes = map[int]string{
	0: "ถ้วน",
	1: "สตางค์",
}

var units = map[int]string{
	0: "ล้าน",
	1: "สิบ",
	2: "ร้อย",
	3: "พัน",
	4: "หมื่น",
	5: "แสน",
}

func getUnitPosition(numLen, pos int) int {
	return (numLen - pos) % len(units)
}

func mapNumToThaiText(numStr string) string {
	var (
		buff bytes.Buffer
		n    = len(numStr)
	)

	for pos, numRune := range numStr {
		num := int(numRune - '0')
		thaiText := digits[num]
		pos = pos + 1
		unitPos := getUnitPosition(n, pos)
		unitText := units[unitPos]

		if num == 0 {
			if unitPos == 0 && pos != n {
				buff.Write([]byte(unitText))
			}
			continue
		}

		buff.Write([]byte(thaiText))

		// for ignore to print unit to last position of number
		if pos != n {
			buff.Write([]byte(unitText))
		}
	}
	return buff.String()
}

// replaceSpecialText replaces special Thai text
func replaceSpecialText(text string) string {
	text = strings.Replace(text, "สิบหนึ่ง", "สิบเอ็ด", -1)
	text = strings.Replace(text, "หนึ่งสิบ", "สิบ", -1)
	text = strings.Replace(text, "สองสิบ", "ยี่สิบ", -1)
	text = strings.Replace(text, "ล้านหนึ่งบาท", "ล้านเอ็ดบาท", -1)
	text = strings.Replace(text, "แสนหนึ่งบาท", "แสนเอ็ดบาท", -1)
	text = strings.Replace(text, "หมื่นหนึ่งบาท", "หมื่นเอ็ดบาท", -1)
	text = strings.Replace(text, "พันหนึ่งบาท", "พันเอ็ดบาท", -1)
	text = strings.Replace(text, "ร้อยหนึ่งบาท", "ร้อยเอ็ดบาท", -1)
	text = strings.Replace(text, "สิบหนึ่งบาท", "สิบเอ็ดบาท", -1)
	return text
}

// convertNumToText return the thai text that already convert number to string.
// This will checks the decimal value in order to insert the currency and suffix
// into the string after convert number to thai text done.
func convertNumToText(numStr string) string {
	var (
		fullText    string
		thaiText    string
		decimalText string
	)

	dotIdx := strings.Index(numStr, ".")
	numWithoutDecimal := numStr[0:dotIdx]
	decimal := numStr[dotIdx+1:]

	thaiText = mapNumToThaiText(numWithoutDecimal)
	if decimal == "00" {
		if thaiText == "" {
			fullText = digits[0] + currency + suffixes[0]
		} else {
			fullText = thaiText + currency + suffixes[0]
		}
	} else {
		decimalText = mapNumToThaiText(decimal)
		if thaiText == "" {
			fullText = decimalText + suffixes[1]
		} else {
			fullText = thaiText + currency + decimalText + suffixes[1]
		}
	}

	return replaceSpecialText(fullText)
}
