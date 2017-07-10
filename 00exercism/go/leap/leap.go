package leap

const testVersion = 3

// IsLeapYear returns true if the int is a leap year
func IsLeapYear(n int) bool {
	// Write some code here to pass the test suite.
	if (n%4 == 0) && ((n%100 != 0) || (n%400 == 0)) {
		return true
	}
	return false
}
