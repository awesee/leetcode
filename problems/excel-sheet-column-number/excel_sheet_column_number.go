package excel_sheet_column_number

func titleToNumber(s string) int {
	ans := 0
	for _, c := range s {
		ans *= 26
		ans += int(c) - 'A' + 1
	}
	return ans
}
