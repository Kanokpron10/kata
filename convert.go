package kata

func convertToDigitalNumber(number int) string {
	if number == 1 {
		return ` 
	|
	|`
	}
	if number == 2 {
		return `_
	_|
	|_`
	}
	if number == 3 {
		return `_
	_|
	_|`
	}
	return ` 
	|_|
	  |`
}
