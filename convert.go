package kata

func convertToDigitalNumber(number int) string {
	digitalNumber := map[int]string{
		1: ` 
	|
	|`,
		2: `_
	_|
	|_`,
		3: `_
	_|
	_|`,
		4: ` 
	|_|
	  |`,
	}
	if number == 1 {
		return digitalNumber[1]
	}
	if number == 2 {
		return digitalNumber[2]
	}
	if number == 3 {
		return digitalNumber[3]
	}
	return digitalNumber[4]
}
