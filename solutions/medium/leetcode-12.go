// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given an integer, convert it to a roman numeral.

// Example 1:

// Input: num = 3
// Output: "III"
// Explanation: 3 is represented as 3 ones.
// Example 2:

// Input: num = 58
// Output: "LVIII"
// Explanation: L = 50, V = 5, III = 3.
// Example 3:

// Input: num = 1994
// Output: "MCMXCIV"
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

// Constraints:

// 1 <= num <= 3999

func intToRoman(num int) string {
	m := make(map[int]string)
	m[1000] = "M"
	m[900] = "CM"
	m[500] = "D"
	m[400] = "CD"
	m[100] = "C"
	m[90] = "XC"
	m[50] = "L"
	m[40] = "XL"
	m[10] = "X"
	m[9] = "IX"
	m[5] = "V"
	m[4] = "IV"
	m[1] = "I"

	var (
		divisor   int
		sortedArr []int
		res       string
	)

	for k, _ := range m {
		sortedArr = append(sortedArr, k)
	}
	sort.Ints(sortedArr)

	for j := len(sortedArr) - 1; j >= 0; j-- {
		divisor = num / sortedArr[j]
		if divisor != 0 {
			for i := 0; i < divisor; i++ {
				res += m[sortedArr[j]]
			}

			num = num % sortedArr[j]
			if num == 0 {
				break
			}
		}

	}

	return res

}