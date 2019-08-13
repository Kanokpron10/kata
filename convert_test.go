package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertToDigitalNumber_By_Number_1_Should_Get_One(t *testing.T) {
	expectedResult := `|
	|`

	actualResult := convertToDigitalNumber(1)

	assert.Equal(t, expectedResult, actualResult, "The two words should be the same.")
}
