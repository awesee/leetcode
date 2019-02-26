package excel_sheet_column_title

func convertToTitle(n int) string {
	if n--; n < 26 {
		return string(n + 'A')
	}
	return convertToTitle(n/26) + string(n%26+'A')
}
