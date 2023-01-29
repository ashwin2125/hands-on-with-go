package conditionals

func In20thCentury(year int) bool {
	if year <= 2000 && year >= 1901 {
		return true
	}
	return false
}
