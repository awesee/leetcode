package excel_sheet_column_title

func convertToTitle(n int) string {
	if n < 27 {
		return string(n - 1 + 'A')
	}
	return convertToTitle((n-1)/26) + string((n-1)%26+'A')
}
