package kata

func convertToDigitalNumber(number int) string {
	if number == 1 {
		return `|
		|`
	}
	if number == 2 {
		return `_
	_|
	|_`
	}
	return `_
	_|
	_|`
}
