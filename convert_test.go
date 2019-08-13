package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertToDigitalNumber_By_Number_1_Should_Get_One(t *testing.T) {
	expectedResult := ` 
	|
	|`

	actualResult := convertToDigitalNumber(1)

	assert.Equal(t, expectedResult, actualResult, "The two words should be the same.")
}

func Test_ConvertToDigitalNumber_By_Number_2_Should_Get_Two(t *testing.T) {
	expectedResult := `_
	_|
	|_`

	actualResult := convertToDigitalNumber(2)

	assert.Equal(t, expectedResult, actualResult, "The two words should be the same.")
}

func Test_ConvertToDigitalNumber_By_Number_3_Should_Get_Three(t *testing.T) {
	expectedResult := `_
	_|
	_|`

	actualResult := convertToDigitalNumber(3)

	assert.Equal(t, expectedResult, actualResult, "The two words should be the same.")
}

func Test_ConvertToDigitalNumber_By_Number_4_Should_Get_Four(t *testing.T) {
	expectedResult := ` 
	|_|
	  |`

	actualResult := convertToDigitalNumber(4)

	assert.Equal(t, expectedResult, actualResult, "The two words should be the same.")
}
