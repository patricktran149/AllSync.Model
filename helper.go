package Model

import (
	"encoding/json"
	"math"
	"reflect"
	"regexp"
	"strings"
)

func RoundNumberDecimalN(number float64, n int) float64 {
	return math.Round(number*math.Pow10(n)) / math.Pow10(n)
}

func EscapeToRegex(input string) string {
	return regexp.QuoteMeta(input)
}

func StringToCompareString(str string) string {
	return strings.ToUpper(strings.TrimSpace(str))
}

func IsItemExistsInArray(item, slice interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		return false //("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func JSONToString(j interface{}) string {
	b, _ := json.Marshal(&j)

	return string(b)
}

func IsStringInArray(str string, arr []string) bool {
	//str = strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(str)), " ", "")
	for _, str2 := range arr {
		//str2 = strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(str2)), " ", "")
		if StringCompare(str, str2) {
			return true
		}
	}
	return false
}

func StringCompare(str1, str2 string) bool {
	str1 = strings.ReplaceAll(strings.ToUpper(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToUpper(str2), " ", "")
	return str1 == str2
}

func IsStringInArrayAndGet(str string, arr []string) (string, bool) {
	str = strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(str)), " ", "")
	for _, str1 := range arr {
		str2 := strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(str1)), " ", "")
		if str == str2 {
			return str1, true
		}
	}
	return "", false
}
